import synthezia_black_logo from "../assets/synthezia-black-logo.png"
import synthezia_white_logo from "../assets/synthezia-white-logo.png"
import synthezia_white_thumb from "../assets/synthezia-white-thumb.png"
import synthezia_black_thumb from "../assets/synthezia-black-thumb.png"
import { useEffect, useState } from "react";
import { useTheme } from "../contexts/ThemeContext";


export function SyntheziaLogo({ className = "", onClick }: { className?: string; onClick?: () => void }) {
  const clickable = typeof onClick === 'function'
  const [isSmallScreen, setIsSmallScreen] = useState(window.innerWidth < 640)
  const { theme } = useTheme()

  useEffect(() => {
    const handleResize = () => {
      setIsSmallScreen(window.innerWidth < 640)
    }

    window.addEventListener('resize', handleResize)
    return () => {
      window.removeEventListener('resize', handleResize)
    }
  }, [])

  const getLogoSrc = () => {
    if (isSmallScreen) {
      return theme === 'dark' ? synthezia_black_thumb : synthezia_white_thumb
    }
    return theme === 'dark' ? synthezia_black_logo : synthezia_white_logo
  }


  return (
    <div className={`${className}`}>
      <img
      // Inclure both light and dark logos for better caching
        src={getLogoSrc()}
        alt="SynthezIA Logo"
        className={`h-8 sm:h-10 w-auto select-none ${clickable ? 'cursor-pointer hover:opacity-90 focus:opacity-90 outline-none' : ''}`}
        role={clickable ? 'button' as const : undefined}
        tabIndex={clickable ? 0 : undefined}
        onClick={onClick}
        onKeyDown={(e) => {
          if (!clickable) return
          if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault()
            onClick?.()
          }
        }}
      />
    </div>
  )
}
