/**
 * @see https://umijs.org/docs/max/access#access
 * */
export default function access(initialState: { currentUser?: API.CurrentUser } | undefined) {
  const { currentUser } = initialState ?? {};
  
  const hasAccess = (accessLevels: string[]) => {
    return accessLevels.includes(currentUser?.access || "");
  };

  return {
    canAdminUser: () => hasAccess(["root", "admin"]),
    canArticle: () => hasAccess(["root", "admin", "manager", "editor"]),
    canArticleCategory: () => hasAccess(["root", "admin", "manager", "editor"]),
    canArticleTag: () => hasAccess(["root", "admin", "manager", "editor"]),
    canWebBanner: () => hasAccess(["root", "admin", "manager"]),
    canWebMeta: () => hasAccess(["root", "admin", "manager"]),
    isRoot: currentUser?.access == "root"
  };
}

export const accessOptions = [
  // {
  //   value: 'root',
  //   label: '超级管理',
  // },
  {
    value: 'admin',
    label: '后台管理',
  },
  {
    value: 'manager',
    label: '网站管理',
  },
  {
    value: 'editor',
    label: '文章编辑',
  },
]