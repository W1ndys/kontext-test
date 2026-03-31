<template>
  <AdminLayout>
    <div class="max-w-3xl">
      <h2 class="text-xl font-bold text-gray-800 mb-6">标签管理</h2>

      <!-- Add Form -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm p-6 mb-6">
        <h3 class="text-sm font-semibold text-gray-700 mb-3">
          {{ editingTag ? '编辑标签' : '添加标签' }}
        </h3>
        <form @submit.prevent="handleSubmit" class="flex gap-3">
          <input
            v-model="formName"
            type="text"
            placeholder="输入标签名称"
            required
            class="flex-1 px-4 py-2 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
          />
          <button
            type="submit"
            :disabled="saving"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-60"
          >
            {{ saving ? '保存中...' : (editingTag ? '更新' : '添加') }}
          </button>
          <button
            v-if="editingTag"
            type="button"
            @click="cancelEdit"
            class="px-4 py-2 bg-gray-100 text-gray-600 text-sm font-medium rounded-lg hover:bg-gray-200 transition-colors"
          >
            取消
          </button>
        </form>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="space-y-3">
        <div v-for="i in 5" :key="i" class="bg-white rounded-xl p-4 animate-pulse flex items-center gap-4">
          <div class="h-5 bg-gray-200 rounded flex-1"></div>
          <div class="h-5 bg-gray-200 rounded w-16"></div>
        </div>
      </div>

      <!-- Tag List -->
      <div v-else class="bg-white rounded-xl border border-gray-100 shadow-sm overflow-hidden">
        <div v-if="tags.length === 0" class="p-8 text-center text-gray-400">
          暂无标签，请在上方添加
        </div>
        <div v-else class="divide-y divide-gray-100">
          <div
            v-for="tag in tags"
            :key="tag.id"
            class="flex items-center justify-between px-6 py-4 hover:bg-gray-50 transition-colors"
          >
            <div class="flex items-center gap-3">
              <span class="inline-flex items-center px-3 py-1 rounded-full bg-blue-50 text-blue-600 text-sm font-medium">
                #{{ tag.name }}
              </span>
              <span class="text-xs text-gray-400">
                {{ tag.article_count || 0 }} 篇文章
              </span>
            </div>
            <div class="flex items-center gap-2">
              <button
                @click="startEdit(tag)"
                class="text-sm text-blue-600 hover:text-blue-700 transition-colors"
              >
                编辑
              </button>
              <button
                @click="handleDelete(tag)"
                class="text-sm text-red-500 hover:text-red-600 transition-colors"
              >
                删除
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center px-4"
      @click.self="showDeleteModal = false"
    >
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-sm">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">确认删除</h3>
        <p class="text-gray-500 text-sm mb-6">
          确定要删除标签「{{ deleteTarget?.name }}」吗？此操作不可撤销。
        </p>
        <div class="flex justify-end gap-3">
          <button
            @click="showDeleteModal = false"
            class="px-4 py-2 text-sm text-gray-600 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
          >
            取消
          </button>
          <button
            @click="confirmDelete"
            :disabled="deleting"
            class="px-4 py-2 text-sm text-white bg-red-500 rounded-lg hover:bg-red-600 transition-colors disabled:opacity-60"
          >
            {{ deleting ? '删除中...' : '确认删除' }}
          </button>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AdminLayout from '@/components/AdminLayout.vue'
import { getTags, createTag, updateTag, deleteTag } from '@/api/tag'

const tags = ref([])
const loading = ref(true)
const formName = ref('')
const editingTag = ref(null)
const saving = ref(false)

const showDeleteModal = ref(false)
const deleteTarget = ref(null)
const deleting = ref(false)

async function fetchTags() {
  loading.value = true
  try {
    const res = await getTags()
    tags.value = res.data || res || []
  } catch (err) {
    console.error('Failed to fetch tags:', err)
  } finally {
    loading.value = false
  }
}

function startEdit(tag) {
  editingTag.value = tag
  formName.value = tag.name
}

function cancelEdit() {
  editingTag.value = null
  formName.value = ''
}

async function handleSubmit() {
  if (!formName.value.trim()) return
  saving.value = true
  try {
    if (editingTag.value) {
      await updateTag(editingTag.value.id, { name: formName.value.trim() })
    } else {
      await createTag({ name: formName.value.trim() })
    }
    formName.value = ''
    editingTag.value = null
    fetchTags()
  } catch (err) {
    console.error('Failed to save tag:', err)
    alert('操作失败：' + (err.response?.data?.error || err.message))
  } finally {
    saving.value = false
  }
}

function handleDelete(tag) {
  deleteTarget.value = tag
  showDeleteModal.value = true
}

async function confirmDelete() {
  if (!deleteTarget.value) return
  deleting.value = true
  try {
    await deleteTag(deleteTarget.value.id)
    showDeleteModal.value = false
    deleteTarget.value = null
    fetchTags()
  } catch (err) {
    console.error('Failed to delete tag:', err)
    alert('删除失败：' + (err.response?.data?.error || err.message))
  } finally {
    deleting.value = false
  }
}

onMounted(() => {
  fetchTags()
})
</script>
