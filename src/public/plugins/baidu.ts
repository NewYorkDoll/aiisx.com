export default () => {
  const hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?5c3fe388523fe948aac9cad50629bc90";
  const s = document.getElementsByTagName("script")[0];
  s.parentNode!.insertBefore(hm, s);
};
