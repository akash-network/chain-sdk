import { describe, expect, it, jest } from "@jest/globals";

import { BatchQueue, type BatchQueueOptions } from "./BatchQueue.ts";

describe(BatchQueue.name, () => {
  describe("add", () => {
    it("adds item to the queue", () => {
      const queue = createBatchQueue();

      queue.add("item1");

      expect(queue.size).toBe(1);
    });

    it("schedules flush on first add", () => {
      const scheduleFn = jest.fn();
      const queue = createBatchQueue({ scheduleFn });

      queue.add("item1");

      expect(scheduleFn).toHaveBeenCalledTimes(1);
      expect(scheduleFn).toHaveBeenCalledWith(expect.any(Function));
    });

    it("does not schedule additional flush when items are added before first flush", () => {
      const scheduleFn = jest.fn();
      const queue = createBatchQueue({ scheduleFn });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");

      expect(scheduleFn).toHaveBeenCalledTimes(1);
    });

    it("schedules new flush after previous flush completes", async () => {
      const scheduleFn = jest.fn((fn: () => void) => queueMicrotask(fn));
      const onFlush = jest.fn();
      const queue = createBatchQueue({ scheduleFn, onFlush });

      queue.add("item1");
      await flushMicrotasks();

      expect(scheduleFn).toHaveBeenCalledTimes(1);

      queue.add("item2");
      expect(scheduleFn).toHaveBeenCalledTimes(2);
    });
  });

  describe("flush behavior", () => {
    it("processes all items when count is less than maxMessagesInTx", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 10,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");

      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(1);
      expect(onFlush).toHaveBeenCalledWith(["item1", "item2", "item3"]);
    });

    it("processes items in batches when count exceeds maxMessagesInTx", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 2,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");
      queue.add("item4");
      queue.add("item5");

      // Need multiple microtask flushes since each batch schedules the next
      await flushMicrotasks();
      await flushMicrotasks();
      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(3);
      expect(onFlush).toHaveBeenNthCalledWith(1, ["item1", "item2"]);
      expect(onFlush).toHaveBeenNthCalledWith(2, ["item3", "item4"]);
      expect(onFlush).toHaveBeenNthCalledWith(3, ["item5"]);
    });

    it("processes exactly maxMessagesInTx items per batch", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 3,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");
      queue.add("item4");
      queue.add("item5");
      queue.add("item6");

      await flushMicrotasks();
      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(2);
      expect(onFlush).toHaveBeenNthCalledWith(1, ["item1", "item2", "item3"]);
      expect(onFlush).toHaveBeenNthCalledWith(2, ["item4", "item5", "item6"]);
    });

    it("schedules next flush when there are more items in queue than `maxMessagesInTx`", async () => {
      const scheduleFn = jest.fn((fn: () => void) => queueMicrotask(fn));
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 1,
        scheduleFn,
      });

      queue.add("item1");
      queue.add("item2");

      // First call schedules flush
      expect(scheduleFn).toHaveBeenCalledTimes(1);

      // First flush processes item1 and schedules next flush for item2
      await flushMicrotasks();
      expect(scheduleFn).toHaveBeenCalledTimes(2);

      // Second flush processes item2
      await flushMicrotasks();
      expect(onFlush).toHaveBeenCalledTimes(2);
    });

    it("does not schedule next flush when queue is emptied", async () => {
      const scheduleFn = jest.fn((fn: () => void) => queueMicrotask(fn));
      const queue = createBatchQueue({
        maxBatchSize: 10,
        scheduleFn,
      });

      queue.add("item1");
      queue.add("item2");

      await flushMicrotasks();
      expect(scheduleFn).toHaveBeenCalledTimes(1);

      await flushMicrotasks();
      expect(scheduleFn).toHaveBeenCalledTimes(1);
    });
  });

  describe("batching", () => {
    it("batches items added in the same microtask", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");

      expect(onFlush).not.toHaveBeenCalled();

      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(1);
      expect(onFlush).toHaveBeenCalledWith(["item1", "item2", "item3"]);
    });

    it("creates separate batches for items added in different microtasks", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");

      await flushMicrotasks();

      queue.add("item3");
      queue.add("item4");

      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(2);
      expect(onFlush).toHaveBeenNthCalledWith(1, ["item1", "item2"]);
      expect(onFlush).toHaveBeenNthCalledWith(2, ["item3", "item4"]);
    });

    it("batches items added before setTimeout fires", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        scheduleFn: (fn) => setTimeout(fn, 0),
      });

      queue.add("item1");
      queue.add("item2");

      expect(onFlush).not.toHaveBeenCalled();

      await delay(10);

      expect(onFlush).toHaveBeenCalledTimes(1);
      expect(onFlush).toHaveBeenCalledWith(["item1", "item2"]);
    });

    it("handles maxMessagesInTx of 1", async () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 1,
        scheduleFn: queueMicrotask,
      });

      queue.add("item1");
      queue.add("item2");
      queue.add("item3");

      await flushMicrotasks();
      await flushMicrotasks();
      await flushMicrotasks();

      expect(onFlush).toHaveBeenCalledTimes(3);
      expect(onFlush).toHaveBeenNthCalledWith(1, ["item1"]);
      expect(onFlush).toHaveBeenNthCalledWith(2, ["item2"]);
      expect(onFlush).toHaveBeenNthCalledWith(3, ["item3"]);
    });

    it("handles synchronous scheduleFn processing items one at a time", () => {
      const onFlush = jest.fn();
      const queue = createBatchQueue({
        onFlush,
        maxBatchSize: 10,
        scheduleFn: (fn) => fn(),
      });

      queue.add("item1");
      expect(onFlush).toHaveBeenCalledTimes(1);
      expect(onFlush).toHaveBeenLastCalledWith(["item1"]);

      queue.add("item2");
      expect(onFlush).toHaveBeenCalledTimes(2);
      expect(onFlush).toHaveBeenLastCalledWith(["item2"]);
    });
  });

  function createBatchQueue<T = string, R = void>(
    overrides?: Partial<BatchQueueOptions<T, R>>,
  ): BatchQueue<T, R> {
    return new BatchQueue<T, R>({
      onFlush: jest.fn() as (items: T[]) => R,
      maxBatchSize: 10,
      scheduleFn: jest.fn(),
      ...overrides,
    });
  }

  function flushMicrotasks(): Promise<void> {
    return new Promise((resolve) => queueMicrotask(resolve));
  }

  function delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
});
