import { Post } from "../posts/post";

export function PostTimeline() {
  return (
    <div>
      {Array.from({ length: 20 }).map((_, i) => (
        // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
        <Post key={i} />
      ))}
    </div>
  );
}