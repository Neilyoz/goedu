module.exports = {
  chainWebpack: config => {
    config.plugin('html').tap((args) => {
      args[0].title = '在线教育系统';
      return args;
    });
  }
}
