import { MetadataRoute } from 'next'
import config from '@/lib/config'

// https://nextjs.org/docs/app/api-reference/file-conventions/metadata/sitemap#generating-a-sitemap-using-code-js-ts
export default function sitemap(): MetadataRoute.Sitemap {
  return [
    {
      url: config.webUrl,
      lastModified: new Date(),
      changeFrequency: 'daily',
      priority: 1,
    }
  ]
}