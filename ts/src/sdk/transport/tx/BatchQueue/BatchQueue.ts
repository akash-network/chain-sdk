export class BatchQueue<T, R> {
  #queue: T[] = [];
  #isScheduledFlush = false;
  readonly #process: (items: T[]) => R;
  readonly #options: BatchQueueOptions<T, R>;

  constructor(options: BatchQueueOptions<T, R>) {
    this.#process = options.onFlush;
    this.#options = options;
  }

  get size(): number {
    return this.#queue.length;
  }

  add(item: T): void {
    this.#queue.push(item);

    if (!this.#isScheduledFlush) {
      this.#isScheduledFlush = true;
      this.#options.scheduleFn(() => {
        this.#flush();
      });
    }
  }

  #flush(): void {
    if (this.#queue.length === 0) {
      this.#isScheduledFlush = false;
      return;
    }

    const itemsToProcess = this.#queue.splice(0, this.#options.maxBatchSize);

    if (this.#queue.length > 0) {
      this.#options.scheduleFn(() => this.#flush());
    } else {
      this.#isScheduledFlush = false;
    }

    this.#process(itemsToProcess);
  }
}

export interface BatchQueueOptions<T, R> {
  onFlush: (items: T[]) => R;
  maxBatchSize: number;
  scheduleFn: (fn: () => void) => void;
}
