/**
 * Test Orchestrator - Manages concurrent test execution with transaction locking
 * 
 * Uses a lock mechanism to ensure blockchain transactions are properly spaced
 * while allowing tests to run concurrently when not creating transactions.
 */

interface LockRequest {
  resolve: () => void;
  timestamp: number;
}

export class TestOrchestrator {
  private static instance: TestOrchestrator;
  private lastTransactionTime: number = 0;
  private readonly minDelayBetweenTransactions: number = 6000; // 12 seconds for account sequence safety
  private lockQueue: LockRequest[] = [];
  private isLocked: boolean = false;
  private lockTimeout?: NodeJS.Timeout;

  private constructor() {}

  static getInstance(): TestOrchestrator {
    if (!TestOrchestrator.instance) {
      TestOrchestrator.instance = new TestOrchestrator();
    }
    return TestOrchestrator.instance;
  }

  /**
   * Wait for the specified time before executing a test
   * @param delayMs - Delay in milliseconds before test execution
   */
  async waitBeforeTest(delayMs: number = 0): Promise<void> {
    if (delayMs > 0) {
      console.log(`Waiting ${delayMs}ms before test execution...`);
      await this.wait(delayMs);
    }
  }

  /**
   * Acquire a lock for transaction execution
   * Ensures minimum delay between transactions while allowing concurrent execution
   */
  async acquireTransactionLock(): Promise<void> {
    return new Promise<void>((resolve) => {
      const request: LockRequest = {
        resolve,
        timestamp: Date.now()
      };

      this.lockQueue.push(request);
      console.log(`Transaction lock: request queued (queue length: ${this.lockQueue.length})`);
      this.processLockQueue();
    });
  }

  /**
   * Release the transaction lock
   * Allows the next queued transaction to proceed
   */
  releaseTransactionLock(): void {
    if (!this.isLocked) {
      console.warn(`Transaction lock: attempted to release lock that wasn't held`);
      return;
    }

    this.lastTransactionTime = Date.now();
    this.isLocked = false;
    
    // Clear the safety timeout
    if (this.lockTimeout) {
      clearTimeout(this.lockTimeout);
      this.lockTimeout = undefined;
    }
    
    console.log(`Transaction lock: released`);
    
    // Process next item in queue immediately - processLockQueue will handle timing
    this.processLockQueue();
  }

  /**
   * Process the lock queue to grant access to the next transaction
   */
  private processLockQueue(): void {
    if (this.isLocked || this.lockQueue.length === 0) {
      return;
    }

    const now = Date.now();
    const timeSinceLastTransaction = now - this.lastTransactionTime;

    // If enough time has passed since last transaction, grant lock immediately
    if (this.lastTransactionTime === 0 || timeSinceLastTransaction >= this.minDelayBetweenTransactions) {
      console.log(`Transaction lock: granting immediately (${timeSinceLastTransaction}ms since last)`);
      this.grantLock();
    } else {
      // Wait for the remaining time, then grant lock
      const waitTime = this.minDelayBetweenTransactions - timeSinceLastTransaction;
      console.log(`Transaction lock: waiting ${waitTime}ms to maintain spacing...`);
      
      setTimeout(() => {
        if (!this.isLocked && this.lockQueue.length > 0) {
          console.log(`Transaction lock: granting after delay`);
          this.grantLock();
        }
      }, waitTime);
    }
  }

  /**
   * Grant the lock to the next request in queue
   */
  private grantLock(): void {
    if (this.lockQueue.length === 0 || this.isLocked) {
      return;
    }

    this.isLocked = true;
    const request = this.lockQueue.shift()!;
    const waitTime = Date.now() - request.timestamp;
    console.log(`Transaction lock: granted to request (waited ${waitTime}ms)`);
    
    // Set a safety timeout to auto-release lock if not released manually
    this.lockTimeout = setTimeout(() => {
      if (this.isLocked) {
        console.warn(`Transaction lock: auto-releasing stuck lock after 30 seconds`);
        this.releaseTransactionLock();
      }
    }, 30000); // 30 second safety timeout
    
    request.resolve();
  }

  /**
   * Reset the orchestrator state (useful for test cleanup)
   */
  reset(): void {
    this.lastTransactionTime = 0;
    this.isLocked = false;
    this.lockQueue = [];
    
    // Clear any pending timeout
    if (this.lockTimeout) {
      clearTimeout(this.lockTimeout);
      this.lockTimeout = undefined;
    }
    
    console.log(`Transaction lock: orchestrator reset`);
  }

  /**
   * Get current lock status (for debugging)
   */
  getLockStatus(): { isLocked: boolean; queueLength: number; timeSinceLastTransaction: number } {
    return {
      isLocked: this.isLocked,
      queueLength: this.lockQueue.length,
      timeSinceLastTransaction: Date.now() - this.lastTransactionTime
    };
  }

  private wait(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}

/**
 * Decorator function to add pre-test delay
 * @param delayMs - Delay in milliseconds before test execution
 */
export function withDelay(delayMs: number) {
  return function(target: any, propertyName: string, descriptor: PropertyDescriptor) {
    const method = descriptor.value;
    descriptor.value = async function(...args: any[]) {
      const orchestrator = TestOrchestrator.getInstance();
      await orchestrator.waitBeforeTest(delayMs);
      return method.apply(this, args);
    };
  };
}

/**
 * Decorator function to ensure transaction spacing using locks
 * Use this on tests that create blockchain transactions
 */
export function withTransactionLock() {
  return function(target: any, propertyName: string, descriptor: PropertyDescriptor) {
    const method = descriptor.value;
    descriptor.value = async function(...args: any[]) {
      const orchestrator = TestOrchestrator.getInstance();
      await orchestrator.acquireTransactionLock();
      try {
        const result = await method.apply(this, args);
        return result;
      } finally {
        orchestrator.releaseTransactionLock();
      }
    };
  };
}

/**
 * Helper functions for use in tests
 */
export const testUtils = {
  /**
   * Wait before test execution
   */
  waitBeforeTest: async (delayMs: number = 0) => {
    const orchestrator = TestOrchestrator.getInstance();
    await orchestrator.waitBeforeTest(delayMs);
  },

  /**
   * Acquire transaction lock (use before creating transactions)
   */
  acquireTransactionLock: async () => {
    const orchestrator = TestOrchestrator.getInstance();
    await orchestrator.acquireTransactionLock();
  },

  /**
   * Release transaction lock (use after transaction completion)
   */
  releaseTransactionLock: () => {
    const orchestrator = TestOrchestrator.getInstance();
    orchestrator.releaseTransactionLock();
  },

  /**
   * Execute a function with transaction lock protection
   */
  withTransactionLock: async <T>(fn: () => Promise<T>): Promise<T> => {
    const orchestrator = TestOrchestrator.getInstance();
    await orchestrator.acquireTransactionLock();
    try {
      return await fn();
    } finally {
      orchestrator.releaseTransactionLock();
    }
  },

  /**
   * Reset orchestrator state
   */
  reset: () => {
    const orchestrator = TestOrchestrator.getInstance();
    orchestrator.reset();
  },

  /**
   * Get lock status for debugging
   */
  getLockStatus: () => {
    const orchestrator = TestOrchestrator.getInstance();
    return orchestrator.getLockStatus();
  },

  /**
   * Simple wait utility
   */
  wait: (ms: number) => new Promise<void>(resolve => setTimeout(resolve, ms))
};
