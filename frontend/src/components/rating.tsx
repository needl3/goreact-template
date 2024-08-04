import IconStar from "./icons/star";

export function Rating({ defaultRating, setRating }: { defaultRating: number, setRating?: (rating: number) => void }) {
  return <div className="flex items-center gap-2">
    {
      Array.from({ length: 5 }).map((_, i) => <IconStar key={i}
        fill={i < defaultRating ? "#f59e0b" : "#94a3b8"}
        onClick={() => setRating && setRating(i+1)}
        className={`${setRating ? "cursor-pointer" : ""}`} />)
    }
  </div>
}
