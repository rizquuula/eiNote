class NotebookPreview {
  ID: string
  Name: string
  IsActive: boolean
  UpdatedAt: Date | null

  constructor(id: string, name: string, updatedAt: Date | null) {
    this.ID = id
    this.Name = name
    this.IsActive = false
    this.UpdatedAt = updatedAt
  }

  SetActive(isActive: boolean): void {
    this.IsActive = isActive
  }
}

export default NotebookPreview