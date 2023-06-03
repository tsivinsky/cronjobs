import axios from "axios";
import { API_URL } from "@/lib/env";

export const $axios = axios.create({
  baseURL: API_URL,
  withCredentials: true,
});
