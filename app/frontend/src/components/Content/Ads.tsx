export default function AdsSection({ children }: { children: React.ReactNode }) {
    return (
        <div className="flex flex-col mt-20">
            <div className=" w-full mx-auto py-4">
                {children}
            </div>
        </div>
    )
}