class NotebookPreview {
  Title: string
  IsActive: boolean

  constructor(title: string) {
    this.Title = title
    this.IsActive = false
  }

  SetActive(isActive: boolean): void {
    this.IsActive = isActive
  }
}

export default NotebookPreview