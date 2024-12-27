'use client'

import { ArrowUp } from "lucide-react"
import { useEffect, useState } from 'react'

// 参考 https://github.com/timlrx/tailwind-nextjs-starter-blog/blob/main/components/ScrollTopAndComment.tsx
export function ScrollTop() {
  const [show, setShow] = useState(false)

  useEffect(() => {
    const handleWindowScroll = () => {
      if (window.scrollY > 50) setShow(true)
      else setShow(false)
    }

    window.addEventListener('scroll', handleWindowScroll)
    return () => window.removeEventListener('scroll', handleWindowScroll)
  }, [])
  const handleScrollTop = () => {
    console.log("asdasd")
    window.scrollTo({ top: 0 })
  }

  return (
    <div
      className={`fixed bottom-8 right-8 sm:bottom-4 sm:right-4 flex-col gap-3 ${show ? 'flex' : 'hidden'}`}
    >
      <button
        aria-label="Scroll To Top"
        onClick={handleScrollTop}
        className="rounded-full bg-gray-200 p-2 text-gray-500 transition-all hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-400 dark:hover:bg-gray-600"
      >
        <ArrowUp className="h-5 w-5" />
      </button>
    </div>
  )
}