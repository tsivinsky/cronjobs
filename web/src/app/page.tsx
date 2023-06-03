import { redirect } from "next/navigation";
import { getServerUser } from "@/api/user";

export default async function Home() {
  const user = await getServerUser();

  if (!user) {
    return redirect("/login");
  }

  return (
    <div>
      <h1>Hello, {user.login}</h1>
    </div>
  );
}
