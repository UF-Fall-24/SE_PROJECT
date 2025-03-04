export const testEnvironment = "jsdom";
export const moduleNameMapper = {
  "^react-router-dom$": "<rootDir>/node_modules/react-router-dom"
};
export const transformIgnorePatterns = ["/node_modules/(?!(react-router-dom)/)"];
export const setupFilesAfterEnv = ["<rootDir>/src/setupTests.js"];
