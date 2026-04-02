export async function pasteImage(file: File): Promise<string> {
  const allowedTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  const maxSize = 5 * 1024 * 1024

  if (!allowedTypes.includes(file.type)) {
    throw new Error('仅支持 JPG、PNG、GIF、WebP 格式图片')
  }

  if (file.size > maxSize) {
    throw new Error('图片大小不能超过 5MB')
  }

  const formData = new FormData()
  formData.append('image', file)

  const response = await fetch('/api/v1/admin/upload', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: formData
  })

  if (!response.ok) {
    throw new Error('图片上传失败')
  }

  const data = await response.json()
  return data.data.url
}

export function handlePaste(event: ClipboardEvent, onImagePasted: (url: string) => void): void {
  const items = event.clipboardData?.items
  if (!items) return

  for (const item of items) {
    if (item.type.startsWith('image/')) {
      event.preventDefault()
      const file = item.getAsFile()
      if (file) {
        pasteImage(file)
          .then(url => onImagePasted(url))
          .catch(err => console.error('Paste error:', err))
      }
      break
    }
  }
}
