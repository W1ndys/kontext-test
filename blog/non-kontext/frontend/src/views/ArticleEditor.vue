<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Edit Article' : 'New Article' }}</h1>
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div>
        <label class="block text-sm text-gray-600 mb-1">Title</label>
        <input v-model="form.title" type="text" class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500" required />
      </div>
      <div class="flex gap-4">
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1">Category</label>
          <input v-model="form.category" type="text" class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500" />
        </div>
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1">Tags (comma separated)</label>
          <input v-model="form.tags" type="text" class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500" placeholder="tag1, tag2" />
        </div>
      </div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Content</label>
        <MdEditor v-model="form.content" language="en-US" :style="{ height: '500px' }" @onUploadImg="handleUploadImg" />
      </div>
      <div class="flex gap-3">
        <button type="submit" :disabled="saving" class="bg-gray-900 text-white px-6 py-2 rounded text-sm hover:bg-gray-800 disabled:opacity-50">
          {{ saving ? 'Saving...' : 'Save' }}
        </button>
        <router-link to="/admin/articles" class="border border-gray-300 px-6 py-2 rounded text-sm text-gray-600 hover:bg-gray-50 no-underline">Cancel</router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getArticle, createArticle, updateArticle, uploadImage } from '../api'

const route = useRoute()
const router = useRouter()
const saving = ref(false)

const isEdit = computed(() => !!route.params.id)

const form = ref({
  title: '',
  content: '',
  category: '',
  tags: '',
})

async function handleUploadImg(files, callback) {
  const urls = []
  for (const file of files) {
    try {
      const { data } = await uploadImage(file)
      urls.push(data.url)
    } catch (e) {
      console.error('Upload failed:', e)
    }
  }
  callback(urls)
}

async function handleSubmit() {
  saving.value = true
  try {
    if (isEdit.value) {
      await updateArticle(route.params.id, form.value)
    } else {
      await createArticle(form.value)
    }
    router.push('/admin/articles')
  } catch (e) {
    alert('Failed to save article')
  } finally {
    saving.value = false
  }
}

onMounted(async () => {
  if (isEdit.value) {
    try {
      const { data } = await getArticle(route.params.id)
      form.value = {
        title: data.title,
        content: data.content,
        category: data.category,
        tags: data.tags,
      }
    } catch (e) {
      alert('Failed to load article')
      router.push('/admin/articles')
    }
  }
})
</script>
