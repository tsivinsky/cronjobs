import { User } from "@/types/user";
import { $axios } from "./axios";
import { cookies } from "next/headers";

export async function getServerUser(): Promise<User | null> {
  const userId = cookies().get("user_id")?.value;
  if (!userId) return null;

  const resp = await $axios.get<User>("/user", {
    headers: {
      Cookie: `user_id=${userId}`,
    },
  });

  return resp.data;
}
