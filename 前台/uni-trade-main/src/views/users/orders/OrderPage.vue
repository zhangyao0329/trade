<!-- 暂时删掉了卖家的评价功能，目前只有买家能评价 -->

<template>
  <UserNav />
  <!-- 我买到的 -->
  <div style="display: flex; justify-content: center; margin: 50px">
    <el-card>
      <el-row style="margin-bottom: 20px; color: dimgray">
        <h3>我买到的</h3>
      </el-row>
      <el-table :data="purchasedData" stripe border empty-text="暂无买到的订单">
        <el-table-column prop="tradeID" label="订单号" width="80"></el-table-column>

        <el-table-column prop="goodsName" label="商品名称" width="150"></el-table-column>
        <el-table-column label="实付" width="80">
          <template #default="{ row }"> {{ row.price + row.shippingCost }}元 </template>
        </el-table-column>
        <el-table-column prop="deliveryMethod" label="发货方式" width="100">
          <template #default="{ row }">
            <span v-if="row.deliveryMethod == '0'">无需快递</span>
            <span v-else-if="row.deliveryMethod == '1'">自提</span>
            <span v-else-if="row.deliveryMethod == '2'">邮寄</span>
            <span v-else>未知方式</span>
          </template>
        </el-table-column>
        <el-table-column prop="shippingAddress.name" label="联系人" width="80"></el-table-column>
        <el-table-column prop="shippingAddress.tel" label="联系电话" width="130"></el-table-column>
        <!-- 用户自己的收货地址 -->
        <el-table-column label="我的地址" width="200">
          <template #default="{ row }">
            {{ row.shippingAddress.province }}
            {{ row.shippingAddress.city }}
            {{ row.shippingAddress.area }}
            {{ row.shippingAddress.detailArea }}
          </template>
        </el-table-column>

        <!-- 查看详情 -->
        <el-table-column label="更多" width="80"
          ><template #default="scope">
            <el-popover effect="light" trigger="hover" placement="top" width="auto">
              <template #default>
                <div>商品金额: {{ scope.row.price }}元</div>
                <div v-if="scope.row.shippingCost != 0">运费: {{ scope.row.shippingCost }}元</div>
                <div>
                  发货地址: {{ scope.row.senderAddress.province }}
                  {{ scope.row.senderAddress.city }}
                  {{ scope.row.senderAddress.area }}
                  {{ scope.row.senderAddress.detailArea }}
                </div>
                <div>卖家: {{ scope.row.sellerName }}</div>
                <div>下单时间: {{ formatTime(scope.row.orderTime) }}</div>
                <div v-if="scope.row.payTime != '0001-01-01T00:00:00Z'">
                  支付时间: {{ formatTime(scope.row.payTime) }}
                </div>
                <div v-if="scope.row.shippingTime != '0001-01-01T00:00:00Z'">
                  发货时间: {{ formatTime(scope.row.shippingTime) }}
                </div>
                <div v-if="scope.row.trackingNumber != ''">快递单号: {{ scope.row.trackingNumber }}</div>
                <div v-if="scope.row.turnoverTime != '0001-01-01T00:00:00Z'">
                  完成时间: {{ formatTime(scope.row.turnoverTime) }}
                </div>
              </template>
              <template #reference>
                <el-button link type="primary" size="small"> 查看详情 </el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>

        <!-- 订单状态 -->
        <el-table-column prop="status" label="订单状态" width="100"></el-table-column>

        <!-- 操作栏 -->
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button
              v-if="scope.row.status == '未发货'"
              :disabled="scope.row.deliveryMethod == '0'"
              size="small"
              type="primary"
              plain
              @click="openAddressEditDialog(scope.row)"
            >
              修改地址
            </el-button>

            <el-button
              v-if="scope.row.status == '已发货'"
              size="small"
              type="primary"
              plain
              @click="handleReceiving(scope.$index, scope.row)"
            >
              确认收货
            </el-button>
            <el-button
              v-if="scope.row.status == '交易完成'"
              size="small"
              type="primary"
              plain
              @click="showCommentDialog(scope.row)"
            >
              去评价
            </el-button>
            <el-button
              v-if="scope.row.status == '未发货' || scope.row.status == '已发货'"
              size="small"
              type="danger"
              @click="showRefundDialog(scope.row)"
            >
              退款
            </el-button>
            <el-button
              v-if="scope.row.status == '退款中'"
              size="small"
              type="primary"
              plain
              @click="handleCancelRefund(scope.$index, scope.row)"
            >
              取消退款
            </el-button>
            <el-button
              v-if="scope.row.status == '未付款'"
              size="small"
              type="primary"
              plain
              @click="handlePay(scope.$index, scope.row)"
            >
              去付款
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div>
        <el-pagination
          size="small"
          style="justify-content: center; margin-top: 20px"
          layout="prev, pager, next"
          :current-page="purchasedPageNum"
          :page-size="purchasedPageSize"
          :total="purchasedTotal"
          @current-change="handlePruchasedPageChange"
        />
      </div>
    </el-card>

    <!-- 切换地址 -->
    <el-dialog v-model="showSelectAddrDialog" title="切换发货地址" width="470px">
      <div class="addressWrapper">
        <div
          class="text item"
          v-for="item in addressData"
          :key="item.id"
          :class="{ active: tempActiveAddress.addrID === item.id }"
          @click="switchAddress(item)"
        >
          <ul>
            <li>
              <span>收<i />货<i />人：</span>{{ item.name }}
            </li>
            <li><span>联系方式：</span>{{ item.tel }}</li>
            <li><span>收货地址：</span>{{ item.province }}{{ item.city }}{{ item.area }}{{ item.detailArea }}</li>
          </ul>
        </div>
      </div>
      <template #footer>
        <span class="dialog-footer" style="display: flex; justify-content: center">
          <el-button @click="cancelSwitchAddress">取消</el-button>
          <el-button type="primary" @click="confirmSwitchAddress">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 评价对话框 -->
    <el-dialog title="评价内容" v-model="commentDialogVisible" width="500px" @close="resetCommentForm">
      <el-input v-model="comment" placeholder="请输入评价内容" :rows="3" type="textarea" style="margin-bottom: 10px">
      </el-input>
      <span class="dialog-footer" style="display: flex; justify-content: center">
        <el-button type="primary" @click="throttledHandleComment">确 定</el-button>
      </span>
    </el-dialog>

    <!-- 退款对话框 -->
    <el-dialog title="退款理由" v-model="refundDialogVisible" width="500px" @close="resetRefundForm">
      <el-input
        v-model="refundReason"
        placeholder="请条理清晰地填写退款理由。纠纷订单将交付管理员处理。"
        :rows="3"
        type="textarea"
        style="margin-bottom: 10px"
      >
      </el-input>
      <span class="dialog-footer" style="display: flex; justify-content: center">
        <el-button type="primary" @click="throttledHandleRefund">确 定</el-button>
      </span>
    </el-dialog>
  </div>

  <!-- 我卖出的 -->
  <div style="display: flex; justify-content: center; margin: 50px">
    <el-card>
      <el-row style="margin-bottom: 20px; color: dimgray"><h3>我卖出的</h3></el-row>
      <el-table :data="selledData" stripe border empty-text="暂无卖出的订单">
        <el-table-column prop="tradeID" label="订单号" width="80"></el-table-column>

        <el-table-column prop="goodsName" label="商品名称" width="150"></el-table-column>
        <el-table-column label="实收" width="80">
          <template #default="{ row }"> {{ row.price + row.shippingCost }}元 </template>
        </el-table-column>
        <el-table-column prop="deliveryMethod" label="发货方式" width="100">
          <template #default="{ row }">
            <span v-if="row.deliveryMethod == '0'">无需快递</span>
            <span v-else-if="row.deliveryMethod == '1'">自提</span>
            <span v-else-if="row.deliveryMethod == '2'">邮寄</span>
            <span v-else>未知方式</span>
          </template>
        </el-table-column>
        <el-table-column prop="shippingAddress.name" label="联系人" width="80"></el-table-column>
        <el-table-column prop="shippingAddress.tel" label="联系电话" width="130"></el-table-column>

        <!-- 买家的收货地址 -->
        <el-table-column label="收货地址" width="200">
          <template #default="{ row }">
            {{ row.shippingAddress.province }}
            {{ row.shippingAddress.city }}
            {{ row.shippingAddress.area }}
            {{ row.shippingAddress.detailArea }}
          </template>
        </el-table-column>

        <!-- 详细信息 -->
        <el-table-column label="更多" width="80"
          ><template #default="scope">
            <el-popover effect="light" trigger="hover" placement="top" width="auto">
              <template #default>
                <div>商品金额: {{ scope.row.price }}元</div>
                <div v-if="scope.row.shippingCost != 0">运费: {{ scope.row.shippingCost }}元</div>
                <div>
                  发货地址: {{ scope.row.senderAddress.province }}
                  {{ scope.row.senderAddress.city }}
                  {{ scope.row.senderAddress.area }}
                  {{ scope.row.senderAddress.detailArea }}
                </div>
                <div>买家: {{ scope.row.sellerName }}</div>
                <div>下单时间: {{ formatTime(scope.row.orderTime) }}</div>
                <div v-if="scope.row.payTime != '0001-01-01T00:00:00Z'">
                  支付时间: {{ formatTime(scope.row.payTime) }}
                </div>
                <div v-if="scope.row.shippingTime != '0001-01-01T00:00:00Z'">
                  发货时间: {{ formatTime(scope.row.shippingTime) }}
                </div>
                <div v-if="scope.row.trackingNumber != ''">快递单号: {{ scope.row.trackingNumber }}</div>
                <div v-if="scope.row.turnoverTime != '0001-01-01T00:00:00Z'">
                  完成时间: {{ formatTime(scope.row.turnoverTime) }}
                </div>
              </template>
              <template #reference>
                <el-button link type="primary" size="small"> 查看详情 </el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>

        <!-- 订单状态 -->
        <el-table-column prop="status" label="订单状态" width="100"></el-table-column>

        <!-- 操作栏 -->
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button
              v-if="scope.row.status == '未发货'"
              size="small"
              type="primary"
              plain
              @click="handleDispatch(scope.row)"
            >
              去发货
            </el-button>

            <el-button
              v-if="scope.row.status == '未发货'"
              size="small"
              type="danger"
              @click="handleCancelOrder(scope.$index, scope.row)"
            >
              取消订单
            </el-button>

            <el-button
              v-if="scope.row.status == '退款中'"
              size="small"
              type="primary"
              plain
              @click="handleAcceptRefund(scope.$index, scope.row)"
            >
              同意退款
            </el-button>
            <el-button
              v-if="scope.row.status == '退款中'"
              size="small"
              type="danger"
              @click="showRejectRefundDialog(scope.row)"
            >
              拒绝退款
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 发货对话框 -->
      <el-dialog title="快递单号" v-model="shipDialogVisible" width="500px">
        <el-input
          v-model="trackingNumber"
          placeholder="请输入快递单号。"
          :rows="3"
          type="textarea"
          style="margin-bottom: 10px"
        >
        </el-input>
        <span class="dialog-footer" style="display: flex; justify-content: center">
          <el-button type="primary" @click="throttledHandleShip">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 拒绝退款对话框 -->
      <el-dialog title="拒绝退款理由" v-model="rejectRefundDialogVisible" width="500px" @close="resetRefundForm">
        <el-input
          v-model="refundReason"
          placeholder="请条理清晰地填写拒绝退款理由。纠纷订单将交付管理员处理。"
          :rows="3"
          type="textarea"
          style="margin-bottom: 10px"
        >
        </el-input>
        <span class="dialog-footer" style="display: flex; justify-content: center">
          <el-button type="primary" @click="throttledHandleRejectRefund">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 分页 -->
      <div>
        <el-pagination
          size="small"
          style="justify-content: center; margin-top: 20px"
          layout="prev, pager, next"
          :current-page="selledPageNum"
          :page-size="selledPageSize"
          :total="selledTotal"
          @current-change="handleSelledPageChange"
        />
      </div>
    </el-card>
  </div>

  <UserFooter />
</template>

<script setup>
import UserNav from '@/components/UserNav.vue'
import UserFooter from '@/components/UserFooter.vue'
import { onMounted, ref, reactive } from 'vue'
import { getPurchasedDataAPI, getSelledDataAPI, operateOrderAPI, editAddressAPI } from '@/api/order.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import useThrottle from '@/hooks/useThrottle'
import { useRouter } from 'vue-router'
import useFormatTime from '@/hooks/useFormatTime'
import { getAddressListAPI } from '@/api/address'

onMounted(() => {
  getAddressList()
})

const showSelectAddrDialog = ref(false)
const addressData = ref([]) // 地址列表
const activeAddress = ref({}) // 激活的地址
const tempActiveAddress = ref({}) // 临时选中的地址
const tempTradeID = ref()
const { formatTime } = useFormatTime()

const { throttled } = useThrottle() // 节流
const router = useRouter()

// 获取地址列表
const getAddressList = async () => {
  const res = await getAddressListAPI()
  addressData.value = res.data.data
  console.log('地址列表：', addressData.value)
}

// 打开切换地址对话框
const openAddressEditDialog = (row) => {
  showSelectAddrDialog.value = true
  activeAddress.value = row.shippingAddress
  console.log('当前地址：', row.shippingAddress)
  tempTradeID.value = row.tradeID
  tempActiveAddress.value = { ...activeAddress.value } // 将当前激活地址存入临时变量
}

// 确认切换地址
const confirmSwitchAddress = async () => {
  activeAddress.value = { ...tempActiveAddress.value } // 更新激活状态
  showSelectAddrDialog.value = false
  const params = {
    id: tempTradeID.value,
    addrID: activeAddress.value.id
  }
  const res = await editAddressAPI(params)
  if (res.data.code === 1) {
    ElMessage.success('修改成功')
    getPurchasedData()
  }
}

// 取消切换地址
const cancelSwitchAddress = () => {
  tempActiveAddress.value = { ...activeAddress.value } // 恢复为之前的激活状态
  showSelectAddrDialog.value = false
}

// 切换地址
const switchAddress = (item) => {
  tempActiveAddress.value = item // 更新临时变量
}

// 从接口拿取“我买到的”订单信息
let purchasedPageNum = ref(1) //表格页码
let purchasedPageSize = ref(5) //每页最大展示条数
const purchasedTotal = ref(0)
const purchasedData = ref([])

// 获取我买到的订单信息
const getPurchasedData = async () => {
  const res = await getPurchasedDataAPI(purchasedPageNum.value, purchasedPageSize.value)
  // console.log('getPurchasedDataAPI响应:', res.data.data)
  purchasedData.value = res.data.data.orderList
  purchasedTotal.value = res.data.data.total
}

// 初始化时获取数据
const handlePruchasedPageChange = (page) => {
  purchasedPageNum.value = page
  getPurchasedData()
}

// 从接口拿取“我卖出的”订单信息
let selledPageNum = ref(1) //表格页码
let selledPageSize = ref(5) //每页最大展示条数
const selledTotal = ref(0)
const selledData = ref([])

// 获取我卖出的订单信息
const getSelledData = async () => {
  const res = await getSelledDataAPI(selledPageNum.value, selledPageSize.value)
  // console.log('getSelledDataAPI响应:', res.data.data)
  selledData.value = res.data.data.orderList
  selledTotal.value = res.data.data.total
}

// 初始化时获取数据
const handleSelledPageChange = (page) => {
  selledPageNum.value = page
  getSelledData()
}

// 对话框状态
const commentDialogVisible = ref(false)
const refundDialogVisible = ref(false)
const rejectRefundDialogVisible = ref(false)
const shipDialogVisible = ref(false)

// 当前操作的订单信息
const currentOrder = reactive({
  id: null,
  targetStatus: ''
})

// 评价和退款的输入内容
const comment = ref('')
const refundReason = ref('')

// 重置评价表单数据
const resetCommentForm = () => {
  comment.value = ''
}

// 重置退款表单数据
const resetRefundForm = () => {
  refundReason.value = ''
}

// 显示评价对话框
const showCommentDialog = (order) => {
  currentOrder.id = order.tradeID
  currentOrder.targetStatus = '已评价'
  commentDialogVisible.value = true
}

// 显示退款对话框
const showRefundDialog = (order) => {
  currentOrder.id = order.tradeID
  currentOrder.targetStatus = '退款中'
  refundDialogVisible.value = true
}

// 显示拒绝退款对话框
const showRejectRefundDialog = (order) => {
  currentOrder.id = order.tradeID
  currentOrder.targetStatus = '处理中'
  rejectRefundDialogVisible.value = true
}

// 支付
const handlePay = async (index, row) => {
  const orderId = row.tradeID
  localStorage.setItem('tradeId', orderId)
  router.push({
    path: '/pay',
    query: {
      id: orderId
    }
  })
}

// 确认评价
const handleComment = async () => {
  if (!comment.value) {
    ElMessage({
      message: '请输入评价内容',
      type: 'warning',
      plain: true
    })
    return
  }

  const res = await operateOrderAPI({
    id: currentOrder.id,
    status: currentOrder.targetStatus,
    comment: comment.value
  })
  if (res.data.code === 1) {
    ElMessage.success('评价成功！')
    // 最新状态
    const currentStatus = res.data.data.status
    // 查找表格中对应的订单并更新状态 插入哪张表？？要不就只有买家能评价吧
    const order = purchasedData.value.find((item) => item.tradeID === currentOrder.id)
    if (order) {
      order.status = currentStatus // 更新表格数据的状态
    }
  } else {
    ElMessage.error('网络请求失败')
  }
  commentDialogVisible.value = false
}
// 节流处理：限制每秒响应一次
const throttledHandleComment = throttled(handleComment, 1000)

// 退款
const handleRefund = async () => {
  if (!refundReason.value) {
    ElMessage({
      message: '请输入退款理由',
      type: 'warning',
      plain: true
    })
    return
  }

  const res = await operateOrderAPI({
    id: currentOrder.id,
    status: currentOrder.targetStatus,
    refundReason: refundReason.value
  })
  if (res.data.code === 1) {
    ElMessage.success('退款申请成功！')
    // 最新状态
    const currentStatus = res.data.data.status
    // 查找表格中对应的订单并更新状态
    const order = purchasedData.value.find((item) => item.tradeID === currentOrder.id)
    if (order) {
      order.status = currentStatus
    }
  } else {
    ElMessage.error('网络请求失败')
  }
  refundDialogVisible.value = false
}
// 节流处理：限制每秒响应一次
const throttledHandleRefund = throttled(handleRefund, 1000)

// 确认收货
const handleReceiving = async (index, row) => {
  try {
    await ElMessageBox.confirm('您确定要确认收货吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await operateOrderAPI({
      id: row.tradeID,
      status: '交易完成'
    })
    if (res.data.code === 1) {
      ElMessage.success('收货成功！')
      const currentStatus = res.data.data.status
      row.status = currentStatus
    } else {
      ElMessage.error('网络请求失败')
    }
  } catch {
    ElMessage.info('操作已取消')
  }
}

// 取消退款
const handleCancelRefund = async (index, row) => {
  try {
    await ElMessageBox.confirm('您确定要取消退款吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await operateOrderAPI({
      id: row.tradeID,
      status: '取消退款'
    })
    if (res.data.code === 1) {
      ElMessage.success('取消退款成功！')
      const currentStatus = res.data.data.status
      row.status = currentStatus
    } else {
      ElMessage.error('网络请求失败')
    }
  } catch {
    ElMessage.info('操作已取消')
  }
}

// 快递单号
const trackingNumber = ref('')
// 存储当前选中的订单行
const currentRow = ref(null)
// 去发货
const handleDispatch = async (row) => {
  // 保存当前的订单行
  currentRow.value = row
  if (row.deliveryMethod === '2') {
    // 邮寄方式才需要填写快递单号
    shipDialogVisible.value = true
    return
  }

  // 自提和无需快递的订单直接发货
  const res = await operateOrderAPI({
    id: row.tradeID,
    status: '已发货'
  })
  if (res.data.code === 1) {
    ElMessage.success('发货成功！')
    const currentStatus = res.data.data.status
    row.status = currentStatus
  } else {
    ElMessage.error('网络请求失败')
  }
}

// 发货操作，只有填写了快递单号才会发请求
const handleShip = async () => {
  // 检查是否填写了快递单号
  if (!trackingNumber.value) {
    ElMessage({
      message: '请填写快递单号',
      type: 'warning',
      plain: true
    })
    return
  }

  try {
    // 发货请求
    const res = await operateOrderAPI({
      id: currentRow.value.tradeID, // 使用传递过来的当前行数据
      status: '已发货',
      trackingNumber: trackingNumber.value // 传递快递单号
    })

    if (res.data.code === 1) {
      ElMessage.success('发货成功！')
      const currentStatus = res.data.data.status
      currentRow.value.status = currentStatus // 更新当前行的状态
      shipDialogVisible.value = false // 发货成功后关闭对话框
      trackingNumber.value = '' // 清空快递单号
    } else {
      ElMessage.error('网络请求失败')
    }
  } catch {
    ElMessage.error('发货失败，发生错误')
  }
}
// 节流处理：限制每秒响应一次
const throttledHandleShip = throttled(handleShip, 1000)

// 取消订单
const handleCancelOrder = async (index, row) => {
  try {
    await ElMessageBox.confirm('您确定要取消订单吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await operateOrderAPI({
      id: row.tradeID,
      status: '已取消'
    })
    if (res.data.code === 1) {
      ElMessage.success('取消订单成功！')
      const currentStatus = res.data.data.status
      row.status = currentStatus
    } else {
      ElMessage.error('网络请求失败')
    }
  } catch {
    ElMessage.info('操作已取消')
  }
}

// 同意退款
const handleAcceptRefund = async (index, row) => {
  try {
    await ElMessageBox.confirm('您确定要同意退款吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await operateOrderAPI({
      id: row.tradeID,
      status: '已退款'
    })
    if (res.data.code === 1) {
      ElMessage.success('已同意退款')
      const currentStatus = res.data.data.status
      row.status = currentStatus
    } else {
      ElMessage.error('网络请求失败')
    }
  } catch {
    ElMessage.info('操作已取消')
  }
}

// 拒绝退款
const handleRejectRefund = async () => {
  if (!refundReason.value) {
    ElMessage({
      message: '请输入退款理由',
      type: 'warning',
      plain: true
    })
    return
  }

  const res = await operateOrderAPI({
    id: currentOrder.id,
    status: '处理中',
    rejectReason: refundReason.value
  })
  if (res.data.code === 1) {
    ElMessage.success('拒绝退款，移交管理员处理')
    // 最新状态
    const currentStatus = res.data.data.status
    // 查找表格中对应的订单并更新状态
    const order = selledData.value.find((item) => item.tradeID === currentOrder.id)
    if (order) {
      order.status = currentStatus
    }
  } else {
    ElMessage.error('网络请求失败')
  }
  rejectRefundDialogVisible.value = false
}
// 节流处理：限制每秒响应一次
const throttledHandleRejectRefund = throttled(handleRejectRefund, 1000)

onMounted(() => {
  getPurchasedData(), getSelledData()
})
</script>

<style scoped lang="scss">
.addressWrapper {
  max-height: 500px;
  overflow-y: auto;
}

.text {
  flex: 1;
  min-height: 90px;
  display: flex;
  align-items: center;

  &.item {
    border: 1px solid #f5f5f5;
    margin-bottom: 10px;
    cursor: pointer;
    border-radius: 5px; // 添加圆角
    transition: border-color 0.3s, background-color 0.3s; // 增加平滑过渡效果

    &.active,
    &:hover {
      border-color: $comColor;
      background: rgba(149, 135, 227, 0.1);
    }

    > ul {
      padding: 10px;
      font-size: 14px;
      line-height: 30px;
    }
  }
}
</style>
