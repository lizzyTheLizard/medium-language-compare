package lizzy.medium.compare.micronaut;

import io.micronaut.data.annotation.MappedEntity;
import io.micronaut.data.annotation.Id;
import java.util.UUID;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@MappedEntity
@Data
@NoArgsConstructor
@AllArgsConstructor
class Issue {
    @Id
    private UUID id;
    private String name;
    private String description;

    void partialUpdate(Issue partialIssue) {
        if (partialIssue.getName() != null) {
            this.name = partialIssue.getName();
        }

        if (partialIssue.getDescription() != null) {
            this.description = partialIssue.getDescription();
        }
    }
}
