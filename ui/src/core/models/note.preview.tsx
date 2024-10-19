class Notepreview {

  Title: string;
  Description: string;
  Timestamp: Date;

  constructor(title: string, description: string, timestamp: Date) {
    this.Title = title
    this.Description = description
    this.Timestamp = timestamp

    this.trimDescription()
  }

  trimDescription() {
    this.Description = this.Description.slice(0, 95) + "..."
  }

  GetDate() {
    return this.Timestamp.getDate() + "/" + this.Timestamp.getMonth()
  }

  GetTime() {
    return this.Timestamp.getHours() + ":" + this.Timestamp.getMinutes() + ":" + this.Timestamp.getSeconds()
  }
}

export default Notepreview