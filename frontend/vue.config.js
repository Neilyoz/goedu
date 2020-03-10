module.exports = {
  chainWebpack: config => {
    config.plugin('html').tap((args) => {
      args[0].title = '随便说';
      return args;
    });
  }
}
