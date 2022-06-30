
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="中文姓名" prop="realName"><el-input
            v-model="queryParams.realName"
            placeholder="请输入中文姓名"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item label="工号" prop="employeeId"><el-input
            v-model="queryParams.employeeId"
            placeholder="请输入工号"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item label="组织架构" prop="deptFullName"><el-input
            v-model="queryParams.deptFullName"
            placeholder="请输入组织架构"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item label="在职状态" prop="status"><el-input
            v-model="queryParams.status"
            placeholder="请输入在职状态"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>

          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:employee:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:employee:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
            >修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:employee:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" :data="employeeList" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" /><el-table-column
            label="中文姓名"
            align="center"
            prop="realName"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="工号"
            align="center"
            prop="employeeId"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="邮箱账号"
            align="center"
            prop="email"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="组织架构"
            align="center"
            prop="deptFullName"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="在职状态"
            align="center"
            prop="status"
            :show-overflow-tooltip="true"
          />
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要修改吗?"
                confirm-button-text="修改"
                @onConfirm="handleUpdate(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:employee:edit']"
                  size="mini"
                  type="text"
                  icon="el-icon-edit"
                >修改
                </el-button>
              </el-popconfirm>
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @onConfirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:employee:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                >删除
                </el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">

            <el-form-item label="组织架构序号" prop="deptId">
              <el-input
                v-model="form.deptId"
                placeholder="组织架构序号"
              />
            </el-form-item>
            <el-form-item label="中文姓名" prop="realName">
              <el-input
                v-model="form.realName"
                placeholder="中文姓名"
              />
            </el-form-item>
            <el-form-item label="工号" prop="employeeId">
              <el-input
                v-model="form.employeeId"
                placeholder="工号"
              />
            </el-form-item>
            <el-form-item label="姓名" prop="pinyin">
              <el-input
                v-model="form.pinyin"
                placeholder="姓名"
              />
            </el-form-item>
            <el-form-item label="" prop="pyAbbr">
              <el-input
                v-model="form.pyAbbr"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="邮箱账号" prop="email">
              <el-input
                v-model="form.email"
                placeholder="邮箱账号"
              />
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <el-input
                v-model="form.phone"
                placeholder="手机号"
              />
            </el-form-item>
            <el-form-item label="组织架构" prop="deptFullName">
              <el-input
                v-model="form.deptFullName"
                placeholder="组织架构"
              />
            </el-form-item>
            <el-form-item label="一级组织架构" prop="deptOneName">
              <el-input
                v-model="form.deptOneName"
                placeholder="一级组织架构"
              />
            </el-form-item>
            <el-form-item label="二级组织架构" prop="deptTwoName">
              <el-input
                v-model="form.deptTwoName"
                placeholder="二级组织架构"
              />
            </el-form-item>
            <el-form-item label="三级组织架构" prop="deptThreeName">
              <el-input
                v-model="form.deptThreeName"
                placeholder="三级组织架构"
              />
            </el-form-item>
            <el-form-item label="四级组织架构" prop="deptFourName">
              <el-input
                v-model="form.deptFourName"
                placeholder="四级组织架构"
              />
            </el-form-item>
            <el-form-item label="五级组织架构" prop="deptFiveName">
              <el-input
                v-model="form.deptFiveName"
                placeholder="五级组织架构"
              />
            </el-form-item>
            <el-form-item label="在职状态" prop="status">
              <el-input
                v-model="form.status"
                placeholder="在职状态"
              />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addEmployee, delEmployee, getEmployee, listEmployee, updateEmployee } from '@/api/admin/employee'

export default {
  name: 'Employee',
  components: {
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      employeeList: [],

      // 关系表类型

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        realName: undefined,
        employeeId: undefined,
        pinyin: undefined,
        deptFullName: undefined,
        status: undefined

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: { realName: [{ required: true, message: '中文姓名不能为空', trigger: 'blur' }],
        employeeId: [{ required: true, message: '工号不能为空', trigger: 'blur' }],
        deptFullName: [{ required: true, message: '组织架构不能为空', trigger: 'blur' }],
        status: [{ required: true, message: '在职状态不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listEmployee(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.employeeList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {

        id: undefined,
        deptId: undefined,
        realName: undefined,
        employeeId: undefined,
        pinyin: undefined,
        pyAbbr: undefined,
        email: undefined,
        phone: undefined,
        deptFullName: undefined,
        deptOneName: undefined,
        deptTwoName: undefined,
        deptThreeName: undefined,
        deptFourName: undefined,
        deptFiveName: undefined,
        status: undefined
      }
      this.resetForm('form')
    },
    getImgList: function() {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function() {
      this.fileOpen = false
    },
    // 关系
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加员工列表'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id =
                row.id || this.ids
      getEmployee(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改员工列表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateEmployee(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addEmployee(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delEmployee({ 'ids': Ids })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {
      })
    }
  }
}
</script>
