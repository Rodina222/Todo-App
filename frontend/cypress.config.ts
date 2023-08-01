import { defineConfig } from 'cypress'
export default {
  component: {
    devServer: {
      framework: 'vue-cli',
      bundler: 'webpack'
    },
    chromeWebSecurity: false
  },

  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    }
  }
}
