import { CustomButton } from "../components/button";
import { routes } from "../routes";

export function Home() {
  return <div className="w-full h-dvh flex items-center justify-center">
    <CustomButton name="Login with Google" link={routes.api.auth.signin} />
  </div>
}
