import { Link } from "react-router-dom"

export function CustomButton({ link, onClick, name, clientLink, customStyle, submit }: { link?: string, onClick?: () => void, name: string, clientLink?: boolean, customStyle?: string, submit?: boolean }) {
  const btnStyle = "bg-primary-800 hover:bg-primary-dark duration-300 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline " + customStyle
  if (link) {
    if (clientLink)
      return <Link to={link} className={btnStyle}>{name}</Link>
    return <a href={link} className={btnStyle}>{name}</a>
  }

  return <button className={btnStyle} onClick={onClick} type={submit ? "submit" : "button"}> {name}</button >
}
