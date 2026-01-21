import type { SDKMethod } from "./client/sdkMetadata.ts";

export type PayloadOf<T extends SDKMethod> = Parameters<T>[0];

export type ResponseOf<T extends SDKMethod> = Awaited<ReturnType<T>>;
