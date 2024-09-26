"use client";

import { Post } from "./post";
import { Tabs } from "./tabs";
import { TimelineTab } from "./timeline-tab";
import { PostForm } from "./post-form";
import { useEffect } from "react";
import { client } from "@/lib/api";

export function Timeline() {
  useEffect(() => {
    (async () => {
      await client.GET("/timeline");
    })();
  }, []);
  return (
    <div className="">
      <div>
        <TimelineTab />
        <div className="h-[68px]"></div>
        <PostForm />
        <div>
          {[...Array(30)].map((_, i) => (
            <Post
              key={i}
              username="kaworu"
              display_name="カヲル"
              avatar_image_url="/kaworu_icon.jpg"
              content="s"
              created_at={1234567}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
