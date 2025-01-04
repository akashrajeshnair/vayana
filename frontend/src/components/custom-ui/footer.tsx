'use client'

export default function Footer() {
  return (
    <footer className="border-t bg-background">
      <div className="container items-center justify-between gap-4 py-5 md:h-8 md:flex-row md:py-0">
        <div className="flex flex-col items-center justify-center gap-4 px-8 md:flex-row md:gap-2 md:px-0">
          <p className="text-center text-sm leading-loose text-muted-foreground">
            Built with love in Bangalore ❤️
          </p>
        </div>
      </div>
    </footer>
  )
}

