import { serve } from "https://deno.land/std/http/server.ts";

const startTime = new Date().getTime();
const host = Deno.env.get("HOSTNAME") ?? Deno.env.get("HOST");
let requests = 0;

function handler(req: Request): Response {

    console.log("Running On:" ,host, "| Total Requests:", ++requests,"| App Uptime:", (new Date().getTime() - startTime)/1000 , "seconds", "| Log Time:",new Date());

    return new Response(`Hello Kubernetes bootcamp deno! | Running on: ${host}  | v=1\n`)
}

serve(handler);
