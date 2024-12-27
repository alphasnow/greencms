import { MetadataRoute } from 'next'
 
// https://nextjs.org/docs/app/api-reference/file-conventions/metadata/robots
export default function robots(): MetadataRoute.Robots {
    // todo: 待优化
  return {
    rules: {
      userAgent: '*',
      allow: '/',
      //disallow: '/private/',
    },
    //sitemap: 'https://acme.com/sitemap.xml',
  }
}