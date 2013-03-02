class CreateModels < ActiveRecord::Migration
  def change
    create_table(:items) do |t|
      t.string :guid
      t.string :some_field
      t.string :some_other_field

      t.timestamps
    end
  end
end
