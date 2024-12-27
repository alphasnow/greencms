export default [
  { path: '/user', layout: false, routes: [{ path: '/user/login', name:"登录", component: './Login' }] },
  { path: '/welcome', name:"首页", icon: 'home', component: './Welcome' },
  {
    path: '/article',
    icon: 'Container',
    name:"文章管理",
    access: 'canArticle',
    routes: [
      // https://pro.ant.design/zh-CN/docs/advanced-menu
      { path: '/article', redirect: '/article/index' },
      { path: '/article/index', name: '文章列表',hideInMenu: true, component: './Article/index' },
      { path: '/article/edit/:id', name: '修改文章',hideInMenu: true,  component: './Article/edit' },
      { path: '/article/create', name: '添加文章',hideInMenu: true,  component: './Article/create'},
    ],
  },
  { path: '/article-category',name:"文章分类", icon: 'Menu', component: './ArticleCategory', access: 'canArticleCategory', },
  { path: '/article-tag',name:"文章标签", icon: 'Function', component: './ArticleTag', access: 'canArticleTag', },
  { path: '/web-meta',name:"网站数据", icon: 'Profile', component: './WebMeta', access: 'canWebMeta', },
  { path: '/web-banner',name:"焦点图", icon: 'FileImage', component: './WebBanner', access: 'canWebBanner', },
  { path: '/admin-user',name:"管理员", icon: 'Safety', component: './AdminUser', access: 'canAdminUser', },
  
  {
    path: '/account',
    icon: 'crown',
    name:"账号管理",
    hideInMenu: true,
    routes: [
      { path: '/account', redirect: '/account/settings' },
      { path: '/account/settings', name: '个人设置', component: './Settings' },
    ],
  },

  { path: '/', redirect: '/welcome' },
  { path: '*', layout: false, component: './404' },
];
