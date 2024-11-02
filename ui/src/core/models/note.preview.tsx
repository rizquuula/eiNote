class NotePreview {

  ID: string
  NotebookID: string
  Title: string
  Content: string
  ContentPreview: string
  Timestamp: Date
  IsActive: boolean

  constructor(id: string, notebookId: string, title: string, content: string, timestamp: Date) {
    this.ID = id
    this.NotebookID = notebookId
    this.Title = title
    this.Content = content
    this.ContentPreview = this.trim(content, 0, 95)
    this.Timestamp = timestamp
    this.IsActive = false
  }

  trim(content: string, start: number, end: number): string {
    if (content.length > end) {
      return content.slice(start, end) + "..."
    } else {
      return content
    }
  }

  SetActive(isActive: boolean): void {
    this.IsActive = isActive
  }

  GetDate(): string {
    return this.Timestamp.getDate() + "/" + this.Timestamp.getMonth()
  }

  GetTime(): string {
    return this.Timestamp.getHours() + ":" + this.Timestamp.getMinutes() + ":" + this.Timestamp.getSeconds()
  }

  UpdateContent(newContent: string): void {
    this.Content = newContent
    this.ContentPreview = this.trim(newContent, 0, 95)
  }
}

export default NotePreview