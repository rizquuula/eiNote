import { THEME_KEY, WebTheme } from "../models/config"

export function SwitchTheme() {
    const currentTheme = localStorage.getItem(THEME_KEY)
    if (currentTheme === WebTheme.DARK) {
        localStorage.setItem(THEME_KEY, WebTheme.LIGHT)
    } else {
        localStorage.setItem(THEME_KEY, WebTheme.DARK)
    }
    window.location.reload()
}