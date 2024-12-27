'use client'

import { MainHeader, MainFooter } from "@/components/main"
import defaultImage from "@/assets/image.jpg"
import Image from "next/image";
import { useSearchParams } from "next/navigation";
import Link from "next/link";
import { Suspense, useEffect, useState } from 'react'

import PageLoading from "@/components/page-loading"
import { getArticleData,getStatistic } from "@/lib/service";
import { Button } from "@/components/ui/button";
import { isoDateToDateTime } from "@/lib/helper";
import { useRouter } from 'next/navigation'
import { useStoreEffect,useActiveStore } from "@/stores";

export default function ArticleShowPage() {
  const searchParams = useSearchParams()
  const id = searchParams.get('id')
  const [article, setArticle] = useState(undefined)
  const router = useRouter()
  const activeStore = useActiveStore()

  const {loaded} = useStoreEffect(async ()=> {
    await getStatistic(id)
    const res = await getArticleData(id)
    setArticle(res)
    document.title = res.title

    activeStore.setCategoryID(res.category_id)
  })
  useEffect(()=>{
    return ()=>{
        activeStore.setCategory("")
    }
  },[])
  if (loaded == false) {
    return <PageLoading />
  }


  return (
      <main className="container px-0 mx-auto mt-3 min-h-screen">

        <div className="flex flex-wrap text-sm">
          <div>
            <Link className="hover:underline" href="/">首页</Link>
          </div>
          <div className="w-6 text-center">/</div>
          <div>
          <Link className="hover:underline" href={`/category?id=${article.article_category.id}`}>{article.article_category.title}</Link>
          </div>
          <div className="w-6 text-center">/</div>
          <div>
            <span className=" text-gray-500">{article.title}</span>
            </div>
        </div>

        <div className="grid gap-3 grid-cols-12 mt-3">
          <div className="lg:col-span-8 sm:col-span-12">
            <div className="bg-white shadow-sm">

              <div className="border-b px-6 py-3 border-gray-100">
                <div className="font-bold text-center text-3xl sm:text-xl">
                  {article.title}
                </div>
                <div className="my-article-content mt-6 border-t min-h-80 text-lg" dangerouslySetInnerHTML={{__html:article.article_content.content}}>
                </div>
                <div className="mt-3 py-3 border-t text-sm text-gray-500">
                <p>特别声明：以上内容(如有图片或视频亦包括在内)均为互联网采集整理，本平台仅提供信息存储服务。</p>
                <p>Notice：The above content (including pictures or videos if there are any) is collected and compiled from the Internet. This platform only provides information storage services.</p>
                </div>
              </div>


            </div>

          </div>
          <div className="lg:col-span-4 sm:col-span-12">
            <div className=" bg-white shadow-sm">

              <img className="w-full min-h-32" src={article.image_url} />

            </div>

            <div className="bg-white shadow-sm mt-3">
              <div className="p-3 min-h-32 text-sm">
                <dl className="flex">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    作者
                  </dt>
                  <dd>
                    <a>{article.origin_author}</a>
                  </dd>
                </dl>

                <dl className="flex mt-3">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    来源
                  </dt>
                  <dd className="text-ellipsis overflow-hidden">
                    <a href={article.origin_url} target="_blank">{article.origin_url}</a>
                  </dd>
                </dl>


                <dl className="flex mt-3">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    分类
                  </dt>
                  <dd>
                    {article.article_category.title}
                  </dd>
                </dl>

                <dl className="flex mt-3">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    标签
                  </dt>
                  {article.article_tags.map((item, key) => {
                    return <dd className="mr-3" key={key}>
                      {item.name}
                    </dd>
                  })}
                </dl>

                
                <dl className="flex mt-3">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    时间
                  </dt>
                  <dd>
                    <a>{isoDateToDateTime(article.created_at)}</a>
                  </dd>
                </dl>

                <dl className="flex mt-3">
                  <dt className="w-12 shrink-0 text-right mr-3">
                    阅读
                  </dt>
                  <dd>
                    <a>{article.article_statistic.views}</a>
                  </dd>
                </dl>

                <div className="flex justify-between text-sm mt-3">
                                    {/* <div>
                    <Button variant="secondary">下一篇</Button>
                  </div> */}
                  <div></div>
                  <div  className="sm:w-full" >
                    <Button className="sm:w-full" size="sm" onClick={()=>router.back()}>返回</Button>
                  </div>

                </div>
                
              </div>
            </div>
          </div>
        </div>

      </main>


  )
}

