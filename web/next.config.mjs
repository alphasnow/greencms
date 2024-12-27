/** @type {import('next').NextConfig} */
const nextConfig = {
    //静态导出 next.config.js
    output: 'export',
    trailingSlash: true,

    // PROD模式下不会出现useEffect执行两次
    // 只是避免在DEV模式下useEffect执行两次
    // https://rishabhsharma.bio/next-js-issue-useeffect-hook-running-twice-in-client-9fb6712f6362
    reactStrictMode: false,

};

export default nextConfig;
