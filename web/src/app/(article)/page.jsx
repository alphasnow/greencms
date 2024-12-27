"use client";

import Image from "next/image";
import { Button } from "@/components/ui/button";
import { ModeToggle } from "@/components/mode-toggle";
import { Badge } from "@/components/ui/badge";
import defaultImage from "@/assets/image.jpg";
import { MainHeader, MainFooter } from "@/components/main";
import config from "@/lib/config";
import ClientInfo from "@/components/client-info";
import { useHomeStore, useStoreEffect } from "@/stores";
import ArticleList from "@/components/article-list";
import { useEffect } from "react";
import Link from "next/link";
import ImageBanner from "@/components/image-banner";
import { useRouter } from 'next/navigation'
import PageLoading from "@/components/page-loading";

export default function HomePage() {
  const homeStore = useHomeStore();
  const router = useRouter()

  const {loaded} = useStoreEffect(async()=>{
    await homeStore.fetchData()
  })
  if(loaded == false){
    return <PageLoading />
  }

  return (
      <main className="container px-0 mx-auto mt-3 min-h-screen">

        <div className="grid gap-3 grid-cols-12">
          <div className="lg:col-span-8 sm:col-span-12">
            <ArticleList list={homeStore.data.articles} />
            <div className="flex justify-center mt-2">
              <Button size="sm" variant="ghost" onClick={() => router.push('/article?page=2')}>点击查看更多文章... </Button>
            </div>

          </div>

          <div className="lg:col-span-4 sm:col-span-12">

            <div className=" bg-white shadow-sm">
              <div className="text-sm border-b p-3 border-gray-100">
                焦点图
              </div>
              <div className="min-h-52">
                <ImageBanner banners={homeStore?.data?.banners || []} />
              </div>
            </div>

            <div className=" bg-white shadow-sm mt-3">
              <div className="text-sm border-b p-3 border-gray-100">
                热门文章
              </div>
              <div className="px-3 min-h-32">
                {homeStore?.data?.hot_articles && (
                  <ul className="pb-3">
                    {homeStore.data.hot_articles.map((item, key) => {
                      return (
                        <li className="my-3" key={key}>
                          <Link
                            className="text-sm font-medium uppercase  hover:text-primary-500 dark:text-gray-300 dark:hover:text-primary-500"
                            href={{
                              pathname: "/article/show",
                              query: { id: item.id },
                            }}
                          >
                            <span>{key+1}. </span>
                            <span className=" hover:text-primary">{item.title}</span>
                            <span className="float-right">阅读量: {item.article_statistic.views}</span>
                            <div className="clear-both"></div>
                          </Link>
                        </li>
                      );
                    })}
                  </ul>
                )}
              </div>
            </div>

            <div className="bg-white shadow-sm mt-3">
              <div className="text-sm border-b p-3 border-gray-100">
                文章标签
              </div>
              <div className="px-3 min-h-32">
                {homeStore?.data?.tags && (
                  <div className="flex flex-wrap">
                    {homeStore.data.tags.map((item, key) => {
                      return (
                        <div className="mb-2 mr-5 mt-2 text-sm" key={key}>
                          <Link
                            className="mr-1 border rounded hover:underline py-1 px-3 text-sm font-medium uppercase text-primary-500 hover:text-primary-600 dark:hover:text-primary-400"
                            style={{color:item.color,borderColor:item.color,}}
                            href={{ pathname: "/tag", query: { id: item.id } }}
                          >
                            {item.name}
                          </Link>
                        </div>
                      );
                    })}
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>
      </main>
  );
}
