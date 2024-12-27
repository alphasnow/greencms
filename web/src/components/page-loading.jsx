import { Skeleton } from "@/components/ui/skeleton"

export default function PageLoading() {
  return (
    <>
    <div className="flex bg-white opacity-95 backdrop-blur-3xl absolute top-0 text-black justify-center items-center h-screen w-screen font-semibold leading-6 text-sm shadow rounded-md  transition ease-in-out duration-150 cursor-not-allowed">
      <svg
        className="animate-spin -ml-1 mr-3 h-5 w-5"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
      >
        <circle
          className="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          strokeWidth="4"
        ></circle>
        <path
          className="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
        ></path>
      </svg>
      数据加载中...
    </div>
    <div className="container px-0 mx-auto mt-3 min-h-screen">
    <Skeleton className="h-4 w-100" />
    <Skeleton className="h-4 mt-4 w-100" />
    <Skeleton className="h-4 mt-4 w-100" />
    <Skeleton className="h-4 mt-4 w-100" />
    <Skeleton className="h-4 mt-4 w-100" />
    </div>
    </>
  );
}
