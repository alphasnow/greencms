"use client";

import { ScrollTop } from "@/components/scroll-top";
import config from "@/lib/config";
import { useGlobalStore } from "@/stores";
import Script from "next/script";

export function MainFooter() {
  const store = useGlobalStore((state) => state);

  let domain = config.webUrl;
  if (domain.includes("127.0.0.1")) {
    domain = "";
  }

  if (!store.isFetched) {
    return null;
  }
  return (
    <footer className="mt-6 py-6 border-t">
      <div className="container mx-auto text-sm text-center ">
      <div>
          <a className="hover:underline" href={config.adminUrl} target="_blank">
            数据管理后台
          </a>

          <span>|</span>

          <a className="hover:underline" href={config.apiUrl} target="_blank">
            服务接口文档
          </a>
        </div>
        <div  className="mt-3">© 2024 {domain}. All rights reserved.</div>
        {store.metas.record_no && (
          <div className="mt-3">ICP备案: {store.metas.record_no}</div>
        )}

      </div>
      <ScrollTop />
      <script
        dangerouslySetInnerHTML={{ __html: store.metas.analytics }}
      ></script>
    </footer>
  );
}
