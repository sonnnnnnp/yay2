import {
  Links,
  Meta,
  type NavigateOptions,
  Outlet,
  Scripts,
  ScrollRestoration,
  isRouteErrorResponse,
  useHref,
  useNavigate,
} from "react-router";

import type { Route } from "./+types/root";
import stylesheet from "./app.css?url";

import { NextUIProvider } from "@nextui-org/system";

import { BackgroundTexture } from "./components/shared/background-texture";

export const links: Route.LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
  { rel: "stylesheet", href: stylesheet },
];

declare module "@react-types/shared" {
  interface RouterConfig {
    routerOptions: NavigateOptions;
  }
}

export default function App() {
  const navigate = useNavigate();

  return (
    <html lang="ja">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="dark font-primary">
        <NextUIProvider navigate={navigate} useHref={useHref}>
          <div className="fixed top-0 left-0 w-full h-dvh overflow-y-scroll text-foreground bg-background">
            <Outlet />
          </div>
          <ScrollRestoration />
          <Scripts />
        </NextUIProvider>
        <BackgroundTexture />
      </body>
    </html>
  );
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  let message = "Oops!";
  let details = "An unexpected error occurred.";
  let stack: string | undefined;

  if (isRouteErrorResponse(error)) {
    message = error.status === 404 ? "404" : "Error";
    details =
      error.status === 404
        ? "The requested page could not be found."
        : error.statusText || details;
  } else if (import.meta.env.DEV && error && error instanceof Error) {
    details = error.message;
    stack = error.stack;
  }

  return (
    <main className="pt-16 p-4 container mx-auto">
      <h1>{message}</h1>
      <p>{details}</p>
      {stack && (
        <pre className="w-full p-4 overflow-x-auto">
          <code>{stack}</code>
        </pre>
      )}
    </main>
  );
}
