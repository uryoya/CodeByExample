module.exports = {
    "testMatch": [
      // "**/__tests__/**/*.+(ts|tsx|js)",
      "**/?(*.)test.(ts|tsx)",
    ],
    "transform": {
      "^.+\\.(ts|tsx)$": "ts-jest"
    },
    moduleDirectories: ['node_modules', 'src'],
  }
