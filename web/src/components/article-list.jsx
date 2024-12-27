"use client";

import Image from "next/image";
import { useEffect } from "react";
import { useHomeStore } from "@/stores";

import { limitString, isoDateToDateTime } from "@/lib/helper";
import { useRouter } from "next/navigation";
import { Button } from "./ui/button";

export default function ArticleList(props) {
  const { list, grid, keywords } = props;
  const router = useRouter();

  const toShow = (article) => {
    router.push("/article/show?id=" + article.id);
  };
  const toTag = (tag) => {
    router.push("/tag?id=" + tag.id);
  };
  const toCategory = (category) => {
    router.push("/category?id=" + category.id);
  };

  if (list.length == 0) {
    return (
      <div className="bg-white shadow-sm min-h-80 flex flex-col justify-center items-center">
        <p>未查询到数据</p>
        <Button
          className="mt-3"
          size="sm"
          variant="secondary"
          onClick={() => router.back()}
        >
          返回
        </Button>
      </div>
    );
  }

  return (
    <div className="bg-white shadow-sm">
      {list.map((item, key) => {
        let formatTitle = item.title;
        if (keywords) {
          formatTitle = formatTitle.replace(
            keywords,
            `<b class="text-red-500">${keywords}</b>`
          );
        }
        let formatDescription = limitString(item.description, 72);
        if (keywords) {
          formatDescription = formatDescription.replace(
            keywords,
            `<b class="text-red-500">${keywords}</b>`
          );
        }
        return (
          <div key={key} className="border-b p-6 border-gray-100">
            <div
              className="text-xl font-bold cursor-pointer hover:text-primary"
              onClick={() => toShow(item)}
              dangerouslySetInnerHTML={{ __html: formatTitle }}
            ></div>
            <div
              className={`grid gap-4 mt-3 ${
                grid == "block" ? "grid-cols-5" : "grid-cols-3"
              }`}
            >
              <div className={(grid == "block" ? "col-span-1" : "col-span-1")+" overflow-hidden rounded"}>
                <img
                  className="w-full h-28 cursor-pointer hover:scale-110 transition duration-300 ease-in-out"
                  onClick={() => toShow(item)}
                  src={item.image_url}
                  alt={item.title}
                />
              </div>
              <div className={grid == "block" ? "col-span-4" : "col-span-2"}>
                <div
                  className="h-20 overflow-hidden indent-8 sm:text-sm"
                  dangerouslySetInnerHTML={{ __html: formatDescription }}
                ></div>
                <div className="h-4 mt-4 flex justify-between text-sm">
                  <div>
                    {/* <span>发布于: {isoDateToDateTime(item.created_at)}</span> */}
                    <span>阅读量:</span> <span>{item.article_statistic.views}</span>
                  </div>
                  <div
                    className="hover:underline cursor-pointer"
                    onClick={() => {
                      toShow(item);
                    }}
                  >
                    阅读原文
                  </div>
                </div>
              </div>
            </div>
            <div className="flex justify-between text-sm mt-3 text-gray-500">
              <div>
                <span>分类:</span>
                <span
                  className="cursor-pointer ml-2  hover:underline"
                  onClick={() => toCategory(item.article_category)}
                >
                  {item.article_category.title}
                </span>
              </div>

              <div>
              <span>标签:</span>
                {item.article_tags.map((i, k) => {
                  return (
                    <span
                      // className={`cursor-pointer mr-2 text-[${i.color}]`}
                      className="cursor-pointer ml-2 hover:underline"
                      // style={{color:i.color}}
                      key={k}
                      onClick={() => toTag(i)}
                    >
                      {i.name}
                    </span>
                  );
                })}
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
