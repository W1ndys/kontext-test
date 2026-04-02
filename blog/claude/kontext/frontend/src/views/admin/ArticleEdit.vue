<template>
  <AdminLayout :title="isEdit ? '编辑文章' : '新建文章'">
    <!-- Loading -->
    <div v-if="pageLoading" class="flex items-center justify-center py-20">
      <svg class="animate-spin h-8 w-8 text-blue-600" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
      </svg>
    </div>

    <div v-else class="space-y-6">
      <!-- Title -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">文章标题</label>
        <input
          v-model="form.title"
          type="text"
          placeholder="请输入文章标题"
          class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
        />
      </div>

      <!-- Category & Tags -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Category -->
        <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">分类</label>
          <select
            v-model="form.category_id"
            class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all bg-white"
          >
            <option :value="0">请选择分类</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">
              {{ cat.name }}
            </option>
          </select>
        </div>

        <!-- Tags -->
        <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">标签</label>
          <div class="flex flex-wrap gap-2">
            <button
              v-for="tag in tags"
              :key="tag.id"
              type="button"
              :class="[
                'px-3 py-1.5 text-sm rounded-full border transition-all',
                selectedTagIds.includes(tag.id)
                  ? 'bg-blue-600 text-white border-blue-600'
                  : 'bg-white text-gray-600 border-gray-300 hover:border-blue-300 hover:text-blue-600'
              ]"
              @click="toggleTag(tag.id)"
            >
              {{ tag.name }}
            </button>
          </div>
          <p v-if="tags.length === 0" class="text-sm text-gray-400 mt-1">暂无标签，请先在标签管理中创建</p>
        </div>
      </div>

      <!-- Markdown Editor -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">文章内容</label>
        <MdEditor
          v-model="form.content"
          :style="{ height: '500px' }"
          language="zh-CN"
          :preview="true"
          @onUploadImg="handleUploadImg"
        />
      </div>

      <!-- Summary -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">
          文章摘要
          <span class="text-gray-400 font-normal">（可选，留空则自动从内容生成）</span>
        </label>
        <textarea
          v-model="form.summary"
          rows="3"
          placeholder="请输入文章摘要..."
          class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all resize-none"
        ></textarea>
      </div>

      <!-- Actions -->
      <div class="flex items-center justify-end gap-3">
        <button
          @click="$router.back()"
          class="px-6 py-2.5 text-sm text-gray-600 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
        >
          取消
        </button>
        <button
          @click="handleSave"
          :disabled="saving"
          class="px-6 py-2.5 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-60 disabled:cursor-not-allowed"
        >
          <span v-if="saving" class="flex items-center gap-2">
            <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            保存中...
          </span>
          <span v-else>{{ isEdit ? '更新文章' : '发布文章' }}</span>
        </button>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import AdminLayout from '@/components/AdminLayout.vue'
import { getArticleDetail, createArticle, updateArticle } from '@/api/article'
import { getCategories } from '@/api/category'
import { getTags } from '@/api/tag'
import { uploadImage } from '@/api/upload'

const route = useRoute()
const router = useRouter()

const isEdit = computed(() => !!route.params.id)
const pageLoading = ref(false)
const saving = ref(false)

const form = reactive({
  title: '',
  content: '',
  summary: '',
  category_id: 0
})

const selectedTagIds = ref([])
const categories = ref([])
const tags = ref([])

function toggleTag(tagId) {
  const idx = selectedTagIds.value.indexOf(tagId)
  if (idx === -1) {
    selectedTagIds.value.push(tagId)
  } else {
    selectedTagIds.value.splice(idx, 1)
  }
}

async function handleUploadImg(files, callback) {
  const results = []
  for (const file of files) {
    try {
      const res = await uploadImage(file)
      const url = res.url || res.data?.url || ''
      results.push({ url, alt: file.name, title: file.name })
    } catch (err) {
      console.error('Upload failed:', err)
    }
  }
  callback(results)
}

async function handleSave() {
  if (!form.title.trim()) {
    alert('请输入文章标题')
    return
  }
  if (!form.content.trim()) {
    alert('请输入文章内容')
    return
  }

  saving.value = true
  try {
    const data = {
      title: form.title,
      content: form.content,
      summary: form.summary || form.content.replace(/[#*`>\-\[\]()!]/g, '').substring(0, 200),
      category_id: form.category_id || undefined,
      tag_ids: selectedTagIds.value
    }

    if (isEdit.value) {
      await updateArticle(route.params.id, data)
    } else {
      await createArticle(data)
    }

    router.push('/admin/articles')
  } catch (err) {
    console.error('Failed to save article:', err)
    alert('保存失败：' + (err.response?.data?.error || err.message))
  } finally {
    saving.value = false
  }
}

onMounted(async () => {
  pageLoading.value = true
  try {
    // Load categories and tags
    const [catRes, tagRes] = await Promise.all([
      getCategories(),
      getTags()
    ])
    categories.value = catRes.data || catRes || []
    tags.value = tagRes.data || tagRes || []

    // Load article data if editing
    if (isEdit.value) {
      const res = await getArticleDetail(route.params.id)
      const article = res.data || res
      form.title = article.title || ''
      form.content = article.content || ''
      form.summary = article.summary || ''
      form.category_id = article.category_id || article.category?.id || 0
      selectedTagIds.value = (article.tags || []).map(t => t.id)
    }
  } catch (err) {
    console.error('Failed to load data:', err)
  } finally {
    pageLoading.value = false
  }
})
</script>
