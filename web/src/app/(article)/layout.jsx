'use client'

import PageLoading from "@/components/page-loading";
import { Suspense } from "react";
import { MainHeader, MainFooter } from "@/components/main";

export default function ArticleLayout({ children }) {
    return (
   
    <Suspense fallback={<PageLoading />}>
    <MainHeader />
        {children}
        <MainFooter />
    </Suspense>

    )
  }