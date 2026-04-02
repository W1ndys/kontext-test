<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { getArticleById, createArticle, updateArticle } from '@/api/article'
import { getCategoryList } from '@/api/category'
import { getTagList } from '@/api/tag'
import { handlePaste } from '@/utils/clipboard'
import ArticleEditor from '@/components/admin/ArticleEditor.vue'
import type { Article, Category, Tag, CreateArticleDTO } from '@/types'

const router = useRouter()
const route = useRoute()
const isEditing = computed(() => !!route.params.id)

const loading = ref(false)
const categories = ref<Category[]>([])
const tags = ref<Tag[]>([])
const form = ref<CreateArticleDTO>({
  title: '',
  slug: '',
  content: '',
  summary: '',
  cover_image: '',
  status: 'draft',
  category_id: 0,
  tag_ids: []
})

onMounted(async () => {
  try {
    const [cats, tgs] = await Promise.all([
      getCategoryList(),
      getTagList()
    ])
    categories.value = cats
    tags.value = tgs

    if (isEditing.value) {
      const article = await getArticleById(Number(route.params.id))
      form.value = {
        title: article.title,
        slug: article.slug,
        content: article.content,
        summary: article.summary,
        cover_image: article.cover_image,
        status: article.status,
        category_id: article.category_id,
        tag_ids: article.tags?.map(t => t.id) || []
      }
    }
  } catch (e) {
    console.error('Failed to load data:', e)
  }
})

async function handleSubmit() {
  if (!form.value.title?.trim()) {
    alert('请输入文章标题')
    return
  }
  if (!form.value.content?.trim()) {
    alert('请输入文章内容')
    return
  }

  loading.value = true
  try {
    if (isEditing.value) {
      await updateArticle(Number(route.params.id), form.value)
    } else {
      await createArticle(form.value)
    }
    router.push('/admin/articles')
  } catch (e) {
    alert('保存失败')
  } finally {
    loading.value = false
  }
}

function onImagePasted(url: string) {
  form.value.content += `\n![image](${url})\n`
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">{{ isEditing ? '编辑文章' : '新建文章' }}</h2>
      <div class="flex gap-3">
        <button @click="router.push('/admin/articles')" class="btn btn-secondary">
          取消
        </button>
        <button @click="handleSubmit" :disabled="loading" class="btn btn-primary">
          {{ loading ? '保存中...' : '保存' }}
        </button>
      </div>
    </div>

    <div class="grid lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2">
        <div class="bg-white rounded-xl shadow-md p-6 mb-6">
          <div class="mb-4">
            <label class="block text-gray-700 mb-2">标题</label>
            <input v-model="form.title" type="text" class="input" placeholder="请输入文章标题" />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 mb-2">别名（URL Slug）</label>
            <input v-model="form.slug" type="text" class="input" placeholder="auto-generated-from-title" />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 mb-2">摘要</label>
            <textarea v-model="form.summary" rows="2" class="input" placeholder="文章摘要（选填）"></textarea>
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 mb-2">封面图片 URL</label>
            <input v-model="form.cover_image" type="text" class="input" placeholder="https://..." />
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-md p-6">
          <label class="block text-gray-700 mb-2">内容（支持 Markdown，粘贴图片自动上传）</label>
          <ArticleEditor
            v-model="form.content"
            @paste="onImagePasted"
          />
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-md p-6 h-fit">
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">状态</label>
          <select v-model="form.status" class="input">
            <option value="draft">草稿</option>
            <option value="published">已发布</option>
          </select>
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">分类</label>
          <select v-model="form.category_id" class="input">
            <option :value="0">请选择分类</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">
              {{ cat.name }}
            </option>
          </select>
        </div>
        <div>
          <label class="block text-gray-700 mb-2">标签</label>
          <div class="space-y-2">
            <label v-for="tag in tags" :key="tag.id" class="flex items-center gap-2">
              <input
                type="checkbox"
                :checked="form.tag_ids?.includes(tag.id)"
                @change="(e: Event) => {
                  const checked = (e.target as HTMLInputElement).checked
                  if (checked) {
                    form.tag_ids = [...(form.tag_ids || []), tag.id]
                  } else {
                    form.tag_ids = (form.tag_ids || []).filter(id => id !== tag.id)
                  }
                }"
                class="rounded text-primary-600"
              />
              <span>{{ tag.name }}</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
