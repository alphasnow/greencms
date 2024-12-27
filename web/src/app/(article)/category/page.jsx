'use client'

import { MainHeader, MainFooter } from "@/components/main"
import defaultImage from "@/assets/image.jpg"
import Image from "next/image";
import { useSearchParams } from "next/navigation";
import Link from "next/link";
import { Suspense, useEffect, useState } from 'react'

import PageLoading from "@/components/page-loading"
import { getArticleList, getCategoryData } from "@/lib/service";
import { Button } from "@/components/ui/button";
import { isoDateToDateTime } from "@/lib/helper";
import { useRouter } from 'next/navigation'
import { useStoreEffect } from "@/stores";
import ArticleList from "@/components/article-list";

export default function CategoryPage() {
  const searchParams = useSearchParams()
  const [data, setData] = useState([])
  const [info, setInfo] = useState()
  const id = searchParams.get("id")
  const router = useRouter();
  const page = searchParams.get("page") || "1";

  const fetchData = async () => {
    let res = null
    res = await getArticleList(`category_id=${id}&page=${page}`)
    setData(res)
    res = await getCategoryData(id)
    setInfo(res)
    document.title = res.title
  }
  const { loaded } = useStoreEffect(async () => {
    await fetchData()
  })
  useEffect(() => {
    if(loaded == false) return 
    fetchData()
  }, [id,page])
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
          {info.title}
          </div>
        </div>

        <div className="mt-3">
        <ArticleList list={data.articles} grid="block" />
        </div>

        {data.articles.length>0 && (
        <div className="bg-white shadow-sm mt-3 p-2 flex justify-between">
          <Button
            className="w-36"
            size="sm"
            disabled={parseInt(page) <= 1}
            onClick={() =>
              router.push(
                `/category?id=${id}&page=${parseInt(page) - 1}`
              )
            }
          >
            上一页
          </Button>
          <Button
            className="w-36"
            size="sm"
            disabled={data.more_articles == false}
            onClick={() =>
              router.push(
                `/category?id=${id}&page=${parseInt(page) + 1}`
              )
            }
          >
            下一页
          </Button>
        </div>
      )}
      </main>

  )
}

