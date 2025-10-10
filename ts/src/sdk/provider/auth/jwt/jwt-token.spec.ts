import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { beforeAll, describe, expect, it } from "@jest/globals";
import fs from "fs";
import path from "path";

import { toBase64Url } from "./base64.ts";
import type { CreateJWTOptions } from "./jwt-token.ts";
import { JwtTokenManager } from "./jwt-token.ts";
import type { ClaimsTestCase, SigningTestCase } from "./test/test-utils.ts";
import { replaceTemplateValues } from "./test/test-utils.ts";
import { createSignArbitraryAkashWallet, type SignArbitraryAkashWallet } from "./wallet-utils.ts";

describe("JWT Claims Validation", () => {
  const testdataPath = path.join(__dirname, "../../../../../..", "testdata", "jwt");
  const jwtMnemonic = fs.readFileSync(path.join(testdataPath, "mnemonic"), "utf-8").trim();
  const jwtSigningTestCases = JSON.parse(fs.readFileSync(path.join(testdataPath, "cases_es256k.json"), "utf-8")) as SigningTestCase[];
  const jwtClaimsTestCases = JSON.parse(fs.readFileSync(path.join(testdataPath, "cases_jwt.json.tmpl"), "utf-8")) as ClaimsTestCase[];

  let testWallet: DirectSecp256k1HdWallet;
  let jwtToken: JwtTokenManager;
  let akashWallet: SignArbitraryAkashWallet;

  beforeAll(async () => {
    testWallet = await DirectSecp256k1HdWallet.fromMnemonic(jwtMnemonic, {
      prefix: "akash",
    });
    akashWallet = await createSignArbitraryAkashWallet(testWallet);
    jwtToken = new JwtTokenManager(akashWallet);
  });

  it.each(jwtClaimsTestCases)("$description", async (testCase) => {
    const { claims, tokenString } = replaceTemplateValues(testCase);

    // For test cases that should fail, we need to validate the payload first
    if (testCase.expected.signFail || testCase.expected.verifyFail) {
      const validationResult = jwtToken.validatePayload(claims);
      expect(validationResult.isValid).toBe(false);

      if (validationResult.isValid) {
        throw new Error("Validation should have failed", { cause: testCase });
      }

      return;
    }

    // For test cases that should pass, create and verify the token
    const token = await jwtToken.generateToken(claims as CreateJWTOptions);
    const decoded = jwtToken.decodeToken(token);
    expect(decoded).toBeDefined();

    // If the test case has a token string, compare it with the generated token
    if (tokenString) {
      expect(token).toEqual(tokenString);
    }
  });

  it.each(jwtSigningTestCases)("$description", async (testCase) => {
    const [expectedHeader, expectedPayload, expectedSignature] = testCase.tokenString.split(".");
    expect(expectedHeader).toBeDefined();
    expect(expectedPayload).toBeDefined();
    expect(expectedSignature).toBeDefined();

    const signingString = `${expectedHeader}.${expectedPayload}`;

    // Sign using the mock wallet's signArbitrary method
    const signResponse = await akashWallet.signArbitrary(akashWallet.address, signingString);
    const signature = toBase64Url(signResponse.signature);

    if (!testCase.mustFail) {
      expect(signature).toBe(expectedSignature);
    } else {
      expect(signature).not.toBe(expectedSignature);
    }
  });
});
