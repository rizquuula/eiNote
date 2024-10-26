class NotebookPreview {
  ID: string
  Title: string
  IsActive: boolean

  constructor(id: string, title: string) {
    this.ID = id
    this.Title = title
    this.IsActive = false
  }

  SetActive(isActive: boolean): void {
    this.IsActive = isActive
  }
}

export default NotebookPreview