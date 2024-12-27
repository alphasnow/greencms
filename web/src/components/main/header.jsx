"use client";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Link from "next/link";

import { useGlobalStore,useActiveStore } from "@/stores";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";
import qs from 'qs';
import {Search,Menu} from 'lucide-react'

export function MainHeader() {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const router = useRouter()
  const [categories, logoImage] = useGlobalStore((state) => [
    state.categories,
    state.metas.logo,
  ]);
  const [keywords,setKeywords] = useState("")
  const activeCategory = useActiveStore((state)=>state.category)
  const [iconShow,setIconShow] = useState("")

  const isActive = (p, s = {}) => {
    if(activeCategory != ""){
      const active = p+"?"+qs.stringify(s)
      return active == activeCategory
    }
    if (p != pathname) return false;
    for (let k in s) {
      if (s[k] != searchParams.get(k)) {
        return false;
      }
    }
    return true;
  };

  const keywordsSearch = () => {
    if(keywords == ""){
      return
    }
    setIconShow("")
    router.push("/search?keywords="+keywords)
  }
  const keywordsEnter = (event) => {
    if (event.keyCode === 13 || event.key === 'Enter') {
      keywordsSearch()
    }
  }
  const toggleIcon = (key) => {
    if(iconShow == key){
      setIconShow("")
      return 
    }
    setIconShow(key)
  }
  return (
    <header className="bg-white shadow-sm ">
      <div className="container px-0 mx-auto flex justify-between h-[3.25rem] items-center">
        <div className="w-24">
        <Link
            href="/"
            onClick={()=>toggleIcon("")}
          >
          <img className="w-4/5" src={logoImage} />
          </Link>
        </div>

        <div className={`lg:grow lg:grid lg:grid-cols-6 lg:gap-1 text-center lg:mx-2 lg:opacity-100 lg:h-auto ${iconShow == "menu" ? "sm:flex sm:flex-col sm:py-2 sm:absolute sm:top-[3.25rem] sm:left-0 sm:mx-0 sm:bg-white sm:border-b sm:border-gray-100 sm:basis-full sm:grid-cols-1 sm:w-full sm:z-20 transition-all sm:duration-300 sm:ease-in-out" : "sm:hidden sm:opacity-0 sm:h-0"}`}>
          <Link
            className={
              (isActive("/") ? "bg-secondary " : "text-gray-500 ") +
              "hover:bg-secondary hover:text-black px-2 py-1 rounded"
            }
            href="/"
            onClick={()=>toggleIcon("")}
          >
            首页
          </Link>
          {categories.map((item, key) => {
            // https://medium.com/little-albie/how-to-pass-query-parameters-from-the-url-in-the-nextjs-14-app-router-ec5690f62c60
            return (
              <Link
                className={
                  (isActive("/category/", { id: item.id })
                    ? "bg-secondary "
                    : "text-gray-500 ") +
                  "hover:bg-secondary hover:text-black px-2 py-1 rounded"
                }
                key={key}
                href={{ pathname: "/category", query: { id: item.id } }}
                onClick={()=>toggleIcon("")}
              >
                {item.title}
              </Link>
            );
          })}
        </div>

        <div className={`lg:w-56 lg:block lg:opacity-100 lg:h-auto ${iconShow == "search" ? "sm:absolute sm:top-[3.25rem] sm:left-0 sm:mx-0 sm:bg-white sm:border-b sm:border-gray-100 sm:basis-full sm:grid-cols-1 sm:w-full sm:z-20 sm:transition-all sm:duration-300 sm:ease-in-out" : "sm:hidden sm:opacity-0 sm:h-0"}`}>
          <div className="flex w-full max-w-sm items-center space-x-2 mx-0 py-0 sm:mx-auto sm:py-2">
            <Input value={keywords} onChange={(e)=>setKeywords(e.target.value)} onKeyUp={keywordsEnter} size="sm" type="text" placeholder="关键词" />
            <Button onClick={keywordsSearch} size="sm" type="button" variants="outline">
              搜索
            </Button>
          </div>
        </div>
        <div className="lg:hidden flex space-x-4">
        <Search onClick={()=>toggleIcon("search")} />
        <Menu onClick={()=>toggleIcon("menu")} />
        </div>
      </div>
      {iconShow != "" && <div onClick={()=>toggleIcon("")} className={`lg:hidden sm:fixed sm:bg-black w-full h-full sm:z-10 transition-all sm:duration-300 sm:ease-in-out sm:opacity-50`}></div> }
    </header>
  );
}
