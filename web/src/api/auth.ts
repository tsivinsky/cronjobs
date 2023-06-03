import { $axios } from "./axios";

export async function logoutUser() {
  const resp = await $axios.get("/auth/logout");

  return resp.data;
}
