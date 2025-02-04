import { Button } from "@/components/ui/button";
import { logout } from "@/store/authSlice";
import { RootState } from "@/store/store";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";


const navigation = [
  { name: 'Home', href: '/' },
  { name: 'About', href: '/about' },
  { name: 'Services', href: '/services' },
  { name: 'Contact', href: '/contact' },
]

export default function Header() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const isAuthenticated = useSelector((state: RootState) => state.auth.isAuthenticated);

  const handleRedirect = (segment: string) => {
    navigate('/' + segment);
  }

  const handleLogout = () => {
    dispatch(logout());
    navigate('/');
  }

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center">
        <div className="mr-4 hidden md:flex">
          <a href="/" className="mx-6 flex items-center space-x-2">
            <span className="hidden font-bold sm:inline-block">Vayana</span>
          </a>
          <nav className="flex items-center space-x-6 text-sm font-medium">
            {navigation.map((item) => (
              <a
                key={item.href}
                href={item.href}
                className={`transition-colors hover:text-foreground/80`}
              >
                {item.name}
              </a>
            ))}
          </nav>
        </div>
        <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
          <div className="w-full flex-1 md:w-auto md:flex-none">
            {/* Add search functionality here if needed */}
          </div>
          {isAuthenticated ? (
            <>
              <nav className="flex items-center">
                <Button variant="ghost" className="ml-2" onClick={() => handleRedirect('login')}>
                  Login
                </Button>
              </nav>
              <nav className="flex items-center">
                <Button variant="ghost" className="ml-2" onClick={() => handleRedirect('register')}>
                  Sign Up
                </Button>
              </nav>
            </>
          ) : (
            <nav className="flex items-center">
              <Button variant="ghost" className="ml-2" onClick={() => handleLogout()}>
                Logout
              </Button>
            </nav>
          )
          }
        </div>
      </div>
    </header>
  )
}

