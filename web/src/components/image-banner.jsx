import * as React from "react";

import { Card, CardContent } from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import Autoplay from "embla-carousel-autoplay";
import Link from "next/link";
import { limitString, isoDateToDateTime } from "@/lib/helper";
// https://ui.shadcn.com/docs/components/carousel

export default function ImageBanner(props) {
  const { banners } = props;
  return (
    <Carousel
      plugins={[
        Autoplay({
          delay: 3000,
        }),
      ]}
      // className="w-full max-w-xs"
    >
      <CarouselContent>
        {banners.map((item, key) => (
          <CarouselItem key={key}>
            <div className="w-full h-full flex justify-center items-center bg-black relative">
              <Link href={item.redirect_url}>
                <img className="w-full" src={item.image_url} />
                {item.description && (
                  <div className="absolute bottom-0 w-full text-white text-center text-sm px-2 py-1 bg-black opacity-80">
                    {limitString(item.description, 20)}
                  </div>
                )}
              </Link>
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
      {/* <CarouselPrevious />
      <CarouselNext /> */}
    </Carousel>
  );
}
