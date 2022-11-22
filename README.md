# Collection of personal dockerfiles

Images are automatically updated by a github workflow which tracks latest release of upstream project
and builds a new image once a new upstream version is available.

CI does not track changes in Dockfiles. If a Dockefile is updated and the image should be rebuilt,
line `if: ${{ steps.exists.outputs.code != 0 }}` in `update_images.yaml` must be commented out.
After image is updated, the line must be uncommented back.