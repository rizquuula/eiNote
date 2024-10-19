class NotePreview {

  Title: string;
  Description: string;
  Timestamp: Date;
  IsActive: boolean;

  constructor(title: string, description: string, timestamp: Date) {
    this.Title = title
    this.Description = description
    this.Timestamp = timestamp
    this.IsActive = false

    this.trimDescription()
  }

  trimDescription(): void {
    this.Description = this.Description.slice(0, 95) + "..."
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
}

export default NotePreview