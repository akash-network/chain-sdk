import type {
  ActionScope as GlobalPermissionScopes,
  JwtTokenPayload,
  LeasePermissionFull,
  LeasePermissionGranular,
  LeasePermissionScoped } from "./validateJwtPayload.ts";

export type AccessScope = GlobalPermissionScopes[number];

export interface JWTHeader {
  alg: string;
  typ: string;
}

export type LeasePermission = LeasePermissionFull | LeasePermissionScoped | LeasePermissionGranular;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type AnyRecord = Record<string, any>;

export type {
  JwtTokenPayload,
};
