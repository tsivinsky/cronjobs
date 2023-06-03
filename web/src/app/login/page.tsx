import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { GitHubButton } from "@/components/GitHubButton";

export default function LoginPage() {
  const c = cookies();
  const userId = c.get("user_id");

  if (userId !== undefined) {
    return redirect("/");
  }

  return (
    <div className="flex justify-center mt-12">
      <GitHubButton />
    </div>
  );
}
