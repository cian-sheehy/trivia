module.exports = {
  lintOnSave: true,
  devServer: {
    port: 80,
    host: "0.0.0.0",
    disableHostCheck: true,
    proxy: {
      "^/api": {
        target:
          "http://" + (process.env.ENV || "trivia-service") + ":3000",
        ws: true,
        changeOrigin: true,
        secure: false,
        logLevel: "debug",
        pathRewrite: { "^/api": "/" },
      }
    },
  },
  transpileDependencies: [
    "vuetify"
  ]
}