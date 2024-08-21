module.exports = {
  extends: [
    'plugin:vue/vue3-essential',
    'eslint:recommended'
  ],
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    ecmaVersion: 2020,
    sourceType: 'module'
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'vue/no-unused-vars': 'warn',
    'vue/no-multiple-template-root': 'off',
    'vue/require-default-prop': 'off',
    'vue/require-prop-types': 'off'
  }
};
